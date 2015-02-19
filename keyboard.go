package main

type Keyboard struct {
	W, A, S, D int
}

func (key *Keyboard) KeyDown(keyCode int) bool {
	update := false
	if keyCode == 87 {
		update = key.W == 0
		key.W = 1
	} else if keyCode == 65 {
		update = key.A == 0
		key.A = 1
	} else if keyCode == 83 {
		update = key.S == 0
		key.S = 1
	} else if keyCode == 68 {
		update = key.D == 0
		key.D = 1
	}
	return update
}

func (key *Keyboard) KeyUp(keyCode int) {
	if keyCode == 87 {
		key.W = 0
	} else if keyCode == 65 {
		key.A = 0
	} else if keyCode == 83 {
		key.S = 0
	} else if keyCode == 68 {
		key.D = 0
	}
}
