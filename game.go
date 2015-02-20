package main

import (
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"time"
)

type Game struct {
	Id      string
	Timer   time.Time
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
	t0 := game.Timer
	game.Timer = time.Now()
	for id, player := range game.Players {
		player.Move(game.Timer.Sub(t0))
		game.Players[id] = player
	}
}

func (game Game) Sleep() {
	s := 1000/60*time.Millisecond - time.Now().Sub(game.Timer)
	if s > 0 {
		time.Sleep(s)
	}
}
