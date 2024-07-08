package Websocket

type OpCode int

const (
	Continuation OpCode = 1
	Text         OpCode = 2
	Binary       OpCode = 3
	Close        OpCode = 8
	Ping         OpCode = 9
	Pong         OpCode = 10
)
