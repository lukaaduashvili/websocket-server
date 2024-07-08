package Websocket

type OpCode int

const (
	Continuation OpCode = 0
	Text         OpCode = 1
	Binary       OpCode = 2
	Close        OpCode = 8
	Ping         OpCode = 9
	Pong         OpCode = 10
)
