package main

import (
	"math/rand"
	"time"
)

type enemy struct {
	x      int
	y      int
	symbol string
}

func newEnemy(x, y int) *enemy {
	return &enemy{x: x, y: y, symbol: "X"}
}

func NewEnemies(n, width, height int) []*enemy {
	var enemies []*enemy
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		enemies = append(enemies, newEnemy(rand.Intn(width), rand.Intn(height)))
	}
	return enemies
}

func (e *enemy) Print(x, y int) (string, bool) {
	if e.x == x && e.y == y {
		return e.symbol, true
	}
	return "", false
}

func PrintEnemies(enemies []*enemy, x, y int) (string, bool) {
	for _, e := range enemies {
		char, ok := e.Print(x, y)
		if ok {
			return char, ok
		}
	}
	return "", false
}
