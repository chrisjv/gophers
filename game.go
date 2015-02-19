package main

import (
	"gopkg.in/igm/sockjs-go.v2/sockjs"
)

type Game struct {
	Id      string
	Players map[string]Player
}

func (game *Game) AddPlayer(session sockjs.Session) {
	player := Player{Session: session}
	game.Players[session.ID()] = player
}

func (game *Game) RemovePlayer(session sockjs.Session) {
	delete(game.Players, session.ID())
}

func (game Game) Sync() {
	msg := Message{Players: game.Players}
	msg.Broadcast()
}

func (game *Game) Play() {
	for id, player := range game.Players {
		player.Move()
		game.Players[id] = player
	}
}
