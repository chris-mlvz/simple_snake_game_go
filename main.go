package main

func main() {
	myGameMap := NewGameMap(widthMap, heightMap, enemiesNum)
	startGame(&myGameMap)
}
