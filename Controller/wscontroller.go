package controller

import (
	"net/http"

	"github.com/gorilla/websocket"
	dataoject "github.com/phy749/LearnEnglish/dataoject"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleWebSocket(hub *dataoject.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	client := &dataoject.Client{
		Hub:  hub,
		Conn: conn,
		Send: make(chan []byte, 256),
	}
	hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}
