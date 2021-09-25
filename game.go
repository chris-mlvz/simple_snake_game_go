package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

const (
	widthMap  = 30
	heightMap = 10
	enemies   = 10
)

func startGame(gameMap *GameMap) {
	var re bool
loop:
	for {
		clear()

		if gameMap.IsGameOver() {
			win()
			break
		}

		fmt.Println("Snake Game")
		gameMap.PrintGame()
		fmt.Println("If you want to move use [w,a,s,d] or")
		fmt.Println("If you want to restart write [r]")
		fmt.Println("If you want to exit write [e]")
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'e':
			close()
			break loop
		case 'r':
			re = true
			break loop
		}

		gameMap.player.Move(char)
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

func win() {
	clear()
	fmt.Println("CONGRATULATIONS, YOU WIN")
}
