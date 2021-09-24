package main

func main() {
	myPlayer := NewPlayer(3, 1)
	myGameMap := NewGameMap(60, 15, &myPlayer)
	startGame(&myPlayer, &myGameMap)
}
