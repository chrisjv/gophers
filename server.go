package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("index.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func StaticFileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func SockJSHandler(session sockjs.Session) {
	fmt.Println("Client connected:", session.ID())
	game.AddPlayer(session)
	msg := Message{Command: "Connect", DataString: session.ID(), Players: game.Players}
	msg.Send(session)
	game.Sync()
	for {
		if msg, err := session.Recv(); err == nil {
			m := Message{}
			json.Unmarshal([]byte(msg), &m)
			if m.Command == "KeyDown" {
				p := game.Players[session.ID()]
				p.Keys.KeyDown(m.DataInt)
				game.Players[session.ID()] = p
			} else if m.Command == "KeyUp" {
				p := game.Players[session.ID()]
				p.Keys.KeyUp(m.DataInt)
				game.Players[session.ID()] = p
			}
			game.Sync()
			continue
		}
		break
	}
	fmt.Println("Client disconnected:", session.ID())
	game.RemovePlayer(session)
	game.Sync()
}

var game = Game{Id: "Game1"}

func main() {
	game.Players = make(map[string]Player)
	go func() {
		for {
			game.Play()
			game.Sleep()
		}
	}()

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/static/", StaticFileHandler)
	http.Handle("/socket/", sockjs.NewHandler("/socket", sockjs.DefaultOptions, SockJSHandler))

	fmt.Println("Starting server...")
	http.ListenAndServe(":8080", nil)
}
