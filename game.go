package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

const (
	widthMap  = 60
	heightMap = 15
	enemies   = 10
)

func startGame(gameMap *GameMap) {
	var re bool
	for {
		clear()

		fmt.Println("Snake Game")
		gameMap.PrintGame()
		fmt.Println("If you want to move use [w,a,s,d]")
		fmt.Println("If you want to restart write [r]")
		fmt.Println("If you want to exit write [e]")
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'e' {
			close()
			break
		}
		if char == 'r' {
			re = true
			break
		}

		gameMap.player.move(char)
		gameMap.player.Kill(gameMap.enemies)
		gameMap.Walls()
	}

	if re {
		restart()
	}
}

func restart() {
	clear()
	myGameMap := NewGameMap(widthMap, heightMap, enemies)
	startGame(&myGameMap)
}

func close() {
	clear()
	fmt.Println("You close the game, good bye")
}
