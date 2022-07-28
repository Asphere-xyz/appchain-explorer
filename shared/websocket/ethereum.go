package websocket

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (s *Server) ethereumWorker(config *Config) error {
	log.Infof("connecting to the go-ethereum websocket url: %s", config.EthereumWebSocketUrl)
	client, err := ethclient.Dial(config.EthereumWebSocketUrl)
	if err != nil {
		return errors.Wrapf(err, "failed to establish connection with go-ethereum")
	}
	headerC := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.TODO(), headerC)
	if err != nil {
		return errors.Wrapf(err, "failed to subscribe to new heads")
	}
	defer sub.Unsubscribe()
	for {
		newHead := <-headerC
		println(newHead.Number.Uint64())
	}
}
