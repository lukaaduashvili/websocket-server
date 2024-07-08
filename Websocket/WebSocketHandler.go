package Websocket

import (
	"fmt"
	"net"
)

type Handler struct {
	openConnections map[string]Websocket
}

func (h *Handler) Initialize() {
	h.openConnections = make(map[string]Websocket)
}

func (h *Handler) HandleTcpConnection(id string, conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing connection")
		}
	}(conn)

	for {
		// Read data from the client
		ws, ok := h.openConnections[id]
		if ok {
			ws.ReceiveMessage()
		} else {
			ws := NewWebsocket(id, conn)
			h.openConnections[id] = *ws
		}
	}
}
