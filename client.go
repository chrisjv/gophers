package main

import (
	"encoding/json"
	"fmt"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/console"
)

var game Game
var sessionId string
var img js.Object

func main() {
	img = js.Global.Get("document").Call("createElement", "img")
	img.Set("src", "/static/gopher.png")

	sockjs := js.Global.Call("SockJS", "http://localhost:8080/socket")

	sockjs.Set("onopen", func() {
		console.Log("Connected to server.")
	})

	sockjs.Set("onmessage", func(e js.Object) {
		msg := Message{}
		json.Unmarshal([]byte(e.Get("data").String()), &msg)
		if msg.Command == "Connect" {
			sessionId = msg.DataString
			console.Log(fmt.Sprintf("Session ID %s", sessionId))
		}
		game.Players = msg.Players
	})

	sockjs.Set("onclose", func() {
		console.Log("Disconnected from server.")
	})

	js.Global.Get("document").Set("onkeydown", func(e js.Object) {
		e.Call("preventDefault")
		me := game.Players[sessionId]
		if me.Keys.KeyDown(e.Get("keyCode").Int()) {
			msg := Message{Command: "KeyDown", DataInt: e.Get("keyCode").Int()}
			msg.Send(sockjs)
		}
		game.Players[sessionId] = me
	})

	js.Global.Get("document").Set("onkeyup", func(e js.Object) {
		e.Call("preventDefault")
		me := game.Players[sessionId]
		me.Keys.KeyUp(e.Get("keyCode").Int())
		msg := Message{Command: "KeyUp", DataInt: e.Get("keyCode").Int()}
		msg.Send(sockjs)
		game.Players[sessionId] = me
	})

	canvas := js.Global.Get("document").Call("getElementById", "canvas")
	ctx := canvas.Call("getContext", "2d")

	render := func() {
		ctx.Call("clearRect", 0, 0, canvas.Get("width"), canvas.Get("height"))
		for id, player := range game.Players {
			player.Move()
			player.Draw(ctx, img)
			game.Players[id] = player
		}
	}

	js.Global.Call("setInterval", render, 1000/60)
}
