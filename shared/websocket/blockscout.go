package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Ankr-network/appchain-explorer/shared/types"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var subExchangeRateNewRate = "[\"6\",\"6\",\"exchange_rate:new_rate\",\"phx_join\",{}]"
var subBlocksIndexing = "[\"3\",\"3\",\"blocks:indexing\",\"phx_join\",{}]"
var subBlocksNewBlock = "[\"12\",\"12\",\"blocks:new_block\",\"phx_join\",{}]"
var subAddressesNewAddress = "[\"9\",\"9\",\"addresses:new_address\",\"phx_join\",{}]"
var subTxsNewTx = "[\"15\",\"15\",\"transactions:new_transaction\",\"phx_join\",{}]"
var subTxsStats = "[\"18\",\"18\",\"transactions:stats\",\"phx_join\",{}]"
var subPhoenixHeatBeat = "[null,\"99\",\"phoenix\",\"heartbeat\",{}]"

func (s *Server) blockscoutWorker(config *Config) error {
	log.Infof("connecting to the blockscout websocket url: %s", config.BlockscoutWebSocketUrl)
	var defaultDialer = &websocket.Dialer{
		Proxy:             http.ProxyFromEnvironment,
		EnableCompression: true,
		HandshakeTimeout:  60 * time.Second,
	}
	con, _, err := defaultDialer.Dial(config.BlockscoutWebSocketUrl, http.Header{})
	if err != nil {
		return errors.Wrapf(err, "failed to connect to the blockscout websocket url")
	}
	closedC := make(chan struct{})
	defer func() {
		closedC <- struct{}{}
		_ = con.Close()
	}()
	// heatbeat worker
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer func() { ticker.Stop() }()
		for {
			select {
			case <-ticker.C:
				_ = con.WriteMessage(websocket.TextMessage, []byte(subPhoenixHeatBeat))
			case <-closedC:
				return
			}
		}
	}()
	// subscribe for channels
	defaultsSubs := []string{
		subExchangeRateNewRate,
		subBlocksIndexing,
		subBlocksNewBlock,
		subAddressesNewAddress,
		subTxsNewTx,
		subTxsStats,
	}
	for _, sub := range defaultsSubs {
		if err := con.WriteMessage(websocket.TextMessage, []byte(sub)); err != nil {
			return errors.Wrapf(err, "failed to subscribe")
		}
	}
	for {
		_, message, err := con.ReadMessage()
		if err != nil {
			return errors.Wrapf(err, "can't read message")
		}
		result, err := parseBlockscoutReply(message)
		if err != nil {
			return errors.Wrapf(err, "can't parse message: %s", string(message))
		}
		switch result.result {
		case "phx_reply":
			// don't print ping/pongs
			if result.channel != "phoenix" {
				log.Infof("subscribed to the blockscout's channel (%s)", result.channel)
			}
			continue
		case "phx_error":
			return fmt.Errorf("handled websocket error: %s", string(message))
		case "new_block":
			block, err := s.blockscoutParseNewBlock(result)
			if err != nil {
				log.WithError(err).Errorf("failed to parse new block")
				continue
			}
			if err := s.broadcastToAll("new_block", block); err != nil {
				log.WithError(err).Errorf("failed to broadcast new block")
				continue
			}
		case "new_transaction":
			transaction, err := s.blockscoutParseNewTransaction(result)
			if err != nil {
				log.WithError(err).Errorf("failed to parse new transaction")
				continue
			}
			if err := s.broadcastToAll("new_transaction", transaction); err != nil {
				log.WithError(err).Errorf("failed to broadcast new transaction")
				continue
			}
		}
	}
}

type blockscoutMessage struct {
	key0, key1, channel, result string
	response                    map[string]interface{}
}

var errBadMessage = fmt.Errorf("bad message")

func parseBlockscoutReply(rawBody []byte) (result blockscoutMessage, err error) {
	var message []interface{}
	if err := json.Unmarshal(rawBody, &message); err != nil {
		return result, errors.Wrapf(err, "failed to marshal response")
	}
	var ok bool
	if result.key0, ok = message[0].(string); message[0] != nil && !ok {
		return result, errBadMessage
	} else if result.key1, ok = message[1].(string); message[1] != nil && !ok {
		return result, errBadMessage
	}
	if result.channel, ok = message[2].(string); !ok {
		return result, errBadMessage
	} else if result.result, ok = message[3].(string); !ok {
		return result, errBadMessage
	}
	result.response, ok = message[4].(map[string]interface{})
	return
}

func (s *Server) blockscoutParseNewBlock(message blockscoutMessage) (*types.BlockDetails, error) {
	var blockNumber float64
	var ok bool
	if blockNumber, ok = message.response["block_number"].(float64); !ok {
		return nil, errBadMessage
	}
	return s.databaseService.GetBlockByNumber(context.TODO(), uint64(blockNumber))
}

func (s *Server) blockscoutParseNewTransaction(message blockscoutMessage) (*types.TransactionDetails, error) {
	var transactionHash string
	var ok bool
	if transactionHash, ok = message.response["transaction_hash"].(string); !ok {
		return nil, errBadMessage
	}
	return s.databaseService.GetTransactionByHash(context.TODO(), transactionHash)
}
