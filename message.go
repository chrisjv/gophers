package main

import (
	"encoding/json"
	"github.com/gopherjs/gopherjs/js"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
)

type Message struct {
	Command    string
	DataString string
	DataInt    int
	Players    map[string]Player
}

func (msg Message) Send(any interface{}) {
	m, _ := json.Marshal(msg)
	switch any.(type) {
	case sockjs.Session:
		session := any.(sockjs.Session)
		session.Send(string(m))
	case js.Object:
		socket := any.(js.Object)
		socket.Call("send", string(m))
	}
}

func (msg Message) Broadcast() {
	m, _ := json.Marshal(msg)
	for _, player := range msg.Players {
		player.Session.Send(string(m))
	}
}
