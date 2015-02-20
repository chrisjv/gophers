package main

import (
	"github.com/gopherjs/gopherjs/js"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"time"
)

type Player struct {
	Session sockjs.Session
	X, Y    int
	Keys    Keyboard
}

func (p *Player) Move(dt time.Duration) {
	p.X += (-p.Keys.A + p.Keys.D) * int(dt/time.Millisecond)
	p.Y += (-p.Keys.W + p.Keys.S) * int(dt/time.Millisecond)
	if p.X < 0 {
		p.X = 0
	} else if p.X > 1024-300 {
		p.X = 1024 - 300
	}
	if p.Y < 0 {
		p.Y = 0
	} else if p.Y > 768-245 {
		p.Y = 768 - 245
	}
}

func (p Player) Draw(ctx js.Object, img js.Object) {
	ctx.Call("drawImage", img, p.X, p.Y, 300, 245)
}
