package main

import "websocket-server/TcpConnectionListener"

func main() {
	// Listen to incoming tcp connections,
	//since that is blocking code run each connection in its own thread
	connectionListener := TcpConnectionListener.TcpConnectionListener{}
	connectionListener.Initialize()
	connectionListener.StartListening()
}
