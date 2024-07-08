package Websocket

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
)

const MAGIC_STRING = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"

func BuildResponse(websocketKey string) http.Response {
	response := http.Response{
		Status:     "101 Switching Protocols",
		StatusCode: 101,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header, 3),
	}

	hasher := sha1.New()
	hasher.Write([]byte(websocketKey))
	hasher.Write([]byte(MAGIC_STRING))
	generatedKey := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	response.Header.Add("Upgrade", "websocket")
	response.Header.Add("Connection", "Upgrade")
	response.Header.Add("Sec-WebSocket-Accept", generatedKey)
	fmt.Printf("%s\n", generatedKey)
	return response
}
