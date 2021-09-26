package main

import (
	"math/rand"
	"time"
)

type Player struct {
	x          int
	y          int
	symbol     rune
	symbolTail rune
	tail       int
	positions  [][2]int
}

func NewPlayer(width, height int) *Player {
	rand.Seed(time.Now().UnixNano())
	return &Player{
		x:          rand.Intn(width),
		y:          rand.Intn(height),
		symbol:     '◉',
		symbolTail: '●',
		tail:       0,
		positions:  [][2]int{},
	}
}

func (p *Player) Move(direction rune) {
	p.controlTail(p.x, p.y)
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

func (p *Player) Set(x, y int) (rune, bool) {
	if p.x == x && p.y == y {
		return p.symbol, true
	}
	return ' ', false
}

func (p *Player) controlTail(x, y int) {
	if p.tail == 0 {
		return
	}
	value := [2]int{x, y}
	index := 0
	if len(p.positions) == index { // nil or empty slice or after last element
		p.positions = append(p.positions, value)
		return
	}
	p.positions = append(p.positions[:index+1], p.positions[index:]...) // index
	p.positions[index] = value
	p.positions = p.positions[:p.tail]
}

func (p *Player) setTail(x, y int) (rune, bool) {
	for _, pos := range p.positions {
		if pos[0] == x && pos[1] == y {
			return p.symbolTail, true
		}
	}
	return ' ', false
}

func (p *Player) Kill(enemies map[[2]int]*enemy) {
	for _, e := range enemies {
		if p.x == e.x && p.y == e.y {
			p.tail++
			delete(enemies, [...]int{p.x, p.y})
		}
	}
}

func (p *Player) checkDeath() bool {
	for _, pos := range p.positions {
		if pos[0] == p.x && pos[1] == p.y {
			return true
		}
	}
	return false
}
