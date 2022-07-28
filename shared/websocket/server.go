package websocket

import (
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
	defer func() { _ = c.Close() }()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
