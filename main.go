package main

import (
	"net/http"
)

func main() {
	SetListenPort(6666)
	SetRequestUri("broadcasting")
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	node := &Node{}
	node.IsAlive = true
	node.Conn = conn

}
