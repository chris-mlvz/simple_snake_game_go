package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

const (
	widthMap   = 10
	heightMap  = 10
	enemiesNum = 10
)

func startGame(gameMap *GameMap) {
	var re bool
loop1:
	for {
		clear()

		if gameMap.checkWin() {
			win()
			break
		}

		if gameMap.player.checkDeath() {
			lose()
			break loop1
		}

		fmt.Println("Snake Game")
		gameMap.PrintGame()
		printInfo(gameMap)

	loop2:
		for {
			char := getChar()
			switch char {
			case 'e':
				close()
				break loop1
			case 'r':
				re = true
				break loop1
			case 'a', 'd', 'w', 's':
				gameMap.player.Move(char)
				gameMap.Walls()
				gameMap.player.Kill(gameMap.enemies)
				gameMap.RespawnEnemies()
				break loop2
			}
		}
	}
	if re {
		restart()
	}
}

func printInfo(gameMap *GameMap) {
	fmt.Println("If you want to move use [w,a,s,d] or")
	fmt.Println("If you want to restart write [r]")
	fmt.Println("If you want to exit write [e]")
	fmt.Printf("You position is: %v,%v\n", gameMap.player.x, gameMap.player.y)
	fmt.Print("Enemies positions: ")
	for _, enemy := range gameMap.enemies {
		fmt.Printf("[%v,%v] ", enemy.x, enemy.y)
	}
	fmt.Println("")
	fmt.Println("Enemy killed: ", gameMap.player.tail)
	fmt.Print("Tail positions: ")
	fmt.Println(gameMap.player.positions)
}

func getChar() rune {
	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		log.Fatal(err)
	}
	return char
}

func restart() {
	clear()
	myGameMap := NewGameMap(widthMap, heightMap, enemiesNum)
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

func lose() {
	clear()
	fmt.Println("You die :(, GAME OVER")
}
