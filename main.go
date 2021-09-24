package main

func main() {
	myGameMap := NewGameMap(widthMap, heightMap, enemies)
	startGame(&myGameMap)
}
