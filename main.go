package main

import (
	"io"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	SetListenPort(8911)
	SetRequestUri("broadcasting")

	http.HandleFunc("/"+requestUri, handler)
	http.ListenAndServe(":"+listenPort, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	node := &Node{}
	node.IsAlive = true
	node.Conn = conn
	node.InStation()

	for {
		mType, reader, err := node.Conn.NextReader()

		if mType == -1 {
			node.Conn.Close()
			node.IsAlive = false
			return
		}

		if err != nil {
			continue
		}

		switch mType {
		case websocket.BinaryMessage:
			continue
		case websocket.CloseMessage:
			node.Conn.Close()
			node.IsAlive = false
			return
		case websocket.PingMessage:
			node.Conn.WriteMessage(websocket.PongMessage, []byte(""))
			continue
		case websocket.PongMessage:
			continue
		}
		//如果是文本内容的话
		//获取文本内容
		msg := make([]byte, 0, 1024)
		for {
			tmp := make([]byte, 1024)
			length, err := reader.Read(tmp)
			if err == io.EOF || length < 1024 {
				msg = append(msg, tmp[:length]...)
				break
			}
			msg = append(msg, tmp...)
		}

		node.PublishMessage(msg)
	}
}
