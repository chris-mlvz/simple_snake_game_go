package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	myPlayer := NewPlayer(3, 1)
	myGameMap := NewGameMap(60, 15, &myPlayer)
	startGame(&myPlayer, &myGameMap)
}

func startGame(player *Player, gameMap *GameMap) {
	for {
		clear()
		fmt.Println("Snake Game")
		gameMap.PrintGame()
		fmt.Println("If you want to move use [w,a,s,d]")
		fmt.Println("If you want to exit write [e]")
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if char == 'e' {
			clear()
			fmt.Println("You close the game, good bye")
			break
		}
		player.move(char)
		// player.Kill(gameMap.enemies)
		gameMap.Walls()
	}
}
