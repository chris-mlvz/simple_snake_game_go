package main

import (
	"math/rand"
	"time"
)

type enemy struct {
	x      int
	y      int
	symbol rune
}

func newEnemy(x, y int) *enemy {
	return &enemy{x: x, y: y, symbol: '✰'}
}

func NewEnemies(n, width, height int, p *Player) map[[2]int]*enemy {
	enemies := make(map[[2]int]*enemy)
	playerPosition := [...]int{p.x, p.y}
	for len(enemies) < n {
		rand.Seed(time.Now().UnixNano())
		w, h := rand.Intn(width), rand.Intn(height)
		newPosition := [...]int{w, h}
		if _, ok := enemies[newPosition]; !ok && newPosition != playerPosition {
			enemies[newPosition] = newEnemy(w, h)
		}
	}
	return enemies
}

func (e *enemy) Set(x, y int) (rune, bool) {
	if e.x == x && e.y == y {
		return e.symbol, true
	}
	return ' ', false
}

func SetEnemies(enemies map[[2]int]*enemy, x, y int) (rune, bool) {
	for _, e := range enemies {
		char, ok := e.Set(x, y)
		if ok {
			return char, ok
		}
	}
	return ' ', false
}
