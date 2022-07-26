package gateway

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type WebSocketServer struct {
}

var webSocketUpgrader = websocket.Upgrader{}

func (s *WebSocketServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := webSocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
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
