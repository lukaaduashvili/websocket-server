package TcpConnectionListener

import (
	"fmt"
	"github.com/google/uuid"
	"net"
	"websocket-server/Websocket"
)

const port = 80

type TcpConnectionListener struct {
	netListener net.Listener
	handler     Websocket.Handler
}

func (listener *TcpConnectionListener) Initialize() {
	listen, err := net.Listen("tcp", ":80")
	listener.netListener = listen
	listener.handler = Websocket.Handler{}
	listener.handler.Initialize()
	if err != nil {
		fmt.Printf("Error initializing tcp listener on port %d: %s\n", port, err)
	}
}

func (listener *TcpConnectionListener) accept() {
	conn, err := listener.netListener.Accept()
	if err != nil {
		fmt.Printf("Error accepting connections: %s\n", err)
		return
	}

	id := uuid.New()

	go listener.handler.HandleTcpConnection(id.String(), conn)
}

func (listener *TcpConnectionListener) StartListening() {
	fmt.Printf("Started listening on port %d\n", port)
	defer func(netListener net.Listener) {
		err := netListener.Close()
		if err != nil {
			fmt.Printf("Error closing listener: %s\n", err)
		}
	}(listener.netListener)

	for {
		listener.accept()
	}
}
