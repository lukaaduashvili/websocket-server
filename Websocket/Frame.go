package Websocket

import "encoding/binary"

type Frame struct {
	IsFragment bool
	Opcode     OpCode
	Reserved   byte
	IsMasked   bool
	Length     uint64
	Mask       []byte
	Payload    []byte
}

func NewFrame(head []byte) *Frame {
	frame := new(Frame)

	frame.IsFragment = head[0]>>7 != 1

	frame.Opcode = OpCode(head[0] & 0x0F)

	maskStart := 2

	if int(head[1])-128 < 126 {
		frame.Length = uint64(int(head[1]) - 128)
	} else if int(head[1]) == 126 {
		length := binary.BigEndian.Uint64(head[2:4])
		frame.Length = length
		maskStart = 4
	} else if int(head[1]) == 127 {
		length := binary.BigEndian.Uint64(head[2:10])
		maskStart = 10
		frame.Length = length
	}

	mask := head[maskStart : maskStart+4]
	frame.Mask = mask

	frame.Payload = head[maskStart+4:]

	return frame
}

func (frame *Frame) GetDecodedPayload() []byte {
	decodedPayload := make([]byte, len(frame.Payload))

	for i := 0; i < len(frame.Payload); i++ {
		decodedPayload[i] = frame.Payload[i] ^ frame.Mask[i&0x3]
	}

	return decodedPayload
}
