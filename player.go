package main

import (
	"math/rand"
	"time"
)

type Player struct {
	x      int
	y      int
	symbol string
}

func NewPlayer(width, height int) *Player {
	rand.Seed(time.Now().UnixNano())
	return &Player{x: rand.Intn(width), y: rand.Intn(height), symbol: "@"}
}

func (p *Player) move(direction rune) {
	switch direction {
	case 'w':
		p.y--
	case 's':
		p.y++
	case 'a':
		p.x--
	case 'd':
		p.x++
	}
}

func (p *Player) Print(x, y int) (string, bool) {
	if p.x == x && p.y == y {
		return p.symbol, true
	}
	return "", false
}

func (p *Player) Kill(enemies []*enemy) {
	for _, e := range enemies {
		if p.x == e.x && p.y == e.y {
			e.symbol = " "
		}
	}
}
