package main

type enemy struct {
	x int
	y int
	symbol string
}

func newEnemy(x, y int) *enemy {
	return &enemy{x: x, y: y, symbol: "X"}
}

func NewEnemies() []*enemy {
	return []*enemy{newEnemy(2, 3), newEnemy(5, 4), newEnemy(3, 4), newEnemy(10, 6)}
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
