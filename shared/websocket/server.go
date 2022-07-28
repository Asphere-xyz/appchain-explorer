package websocket

import (
	"encoding/json"
	"github.com/Ankr-network/ankr-protocol/shared"
	"github.com/Ankr-network/ankr-protocol/shared/database"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Server struct {
	databaseService *database.Service
	// state
	upgrader websocket.Upgrader
	conns    []*websocket.Conn
}

func NewServer(databaseService *database.Service) *Server {
	return &Server{databaseService: databaseService}
}

func (s *Server) Start(cp shared.IConfigProvider) error {
	config := &Config{}
	if err := cp.Parse(config); err != nil {
		return err
	}
	go func() {
		for {
			if err := s.blockscoutWorker(config); err != nil {
				log.WithError(err).Errorf("can't connect to blockscout")
			}
			time.Sleep(10 * time.Second)
		}
	}()
	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Debugf("failed to establish websocket upgrade: %v", err)
		return
	}
	s.conns = append(s.conns, c)
}

func (s *Server) broadcastToAll(channel string, data interface{}) error {
	type messagePayload struct {
		Channel string      `json:"channel"`
		Data    interface{} `json:"data"`
	}
	jsonMessage, err := json.Marshal(messagePayload{Channel: channel, Data: data})
	if err != nil {
		return err
	}
	for _, c := range s.conns {
		_ = c.WriteMessage(websocket.TextMessage, jsonMessage)
	}
	return nil
}
