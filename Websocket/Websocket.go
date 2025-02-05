package Websocket

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"net/http"
)

// Websocket TODO[LA] Handle full duplex communication and persist websockets
type Websocket struct {
	id         string
	conn       net.Conn
	buffer     []byte
	transcript []Frame
}

func NewWebsocket(id string, conn net.Conn) *Websocket {
	ws := Websocket{
		id:         id,
		conn:       conn,
		buffer:     make([]byte, 1024),
		transcript: make([]Frame, 32),
	}

	n, err := ws.conn.Read(ws.buffer)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	reader := bufio.NewReader(bytes.NewReader(ws.buffer[:n]))
	httpRequest, err := http.ReadRequest(reader)

	if err != nil {
		fmt.Println("Error parsing connection headers:", err)
		return nil
	}

	websocketKey := httpRequest.Header.Get("Sec-WebSocket-Key")

	// id value passed in through connection header to handle communication disconnect/reconnect
	websocketReconnectId := httpRequest.Header.Get("connectionId")

	fmt.Printf("Connection Id: %s \n", websocketReconnectId)

	response := BuildResponse(websocketKey)
	err = response.Write(conn)
	if err != nil {
		fmt.Println("Error writing response:", err)
		return nil
	}
	// Process and use the data (here, we'll just print it)
	fmt.Printf("Ack websocket created with key: %s\n", websocketKey)
	return &ws
}

func (ws *Websocket) ReceiveMessage() {
	n, err := ws.conn.Read(ws.buffer)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	frame := NewFrame(ws.buffer[:n])
	ws.transcript = append(ws.transcript, *frame)
	fmt.Printf("Message received: %s \n", frame.GetDecodedPayload())
}
