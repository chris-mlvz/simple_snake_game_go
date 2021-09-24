package main

import (
	"fmt"
	"strings"
)

type GameMap struct {
	width   int
	height  int
	player  *Player
	enemies []*enemy
}

func NewGameMap(width, height, enemyNumber int) GameMap {
	return GameMap{width: width,
		height:  height,
		player:  NewPlayer(width, height),
		enemies: NewEnemies(enemyNumber, width, height)}
}

func (m *GameMap) PrintGame() {
	m.printHLine()
	for y := 0; y < m.height; y++ {
		m.printStartLine()
		for x := 0; x < m.width; x++ {
			m.printChar(x, y)
		}
		m.printEndLine()
	}
	m.printHLine()
}

func (m *GameMap) printChar(x, y int) {
	char_to_draw := " "
	char, ok := PrintEnemies(m.enemies, x, y)
	if ok {
		char_to_draw = char
	}
	char, ok = m.player.Print(x, y)
	if ok {
		char_to_draw = char
	}
	fmt.Print(char_to_draw)
}

func (m *GameMap) printHLine() {
	fmt.Printf("+%v+\n", strings.Repeat("-", m.width))
}

func (m *GameMap) printStartLine() {
	fmt.Print("|")
}

func (m *GameMap) printEndLine() {
	fmt.Println("|")
}

func (m *GameMap) Walls() {
	if m.player.x > m.width-1 {
		m.player.x = 0
	} else if m.player.x < 0 {
		m.player.x = m.width - 1
	}

	if m.player.y > m.height-1 {
		m.player.y = 0
	} else if m.player.y < 0 {
		m.player.y = m.height - 1
	}
}
