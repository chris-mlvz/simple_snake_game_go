package main

type Player struct {
	x int
	y int
}

func NewPlayer(x, y int) Player {
	return Player{x: x, y: y}
}

func (p *Player) move(direction rune) {
	switch direction {
	case 'w':
		p.y--
	case 's':
		p.y++
	case 'a':
		p.x--
	case 'd':
		p.x++
	}
}

func (p *Player) Print(x, y int) (string, bool) {
	if p.x == x && p.y == y {
		return "@", true
	}
	return "", false
}

// func (p *Player) Kill(enemies []*enemy) {
// 	for i := 0; i < len(enemies); i++ {
// 		if p.x == enemies[i].x && p.y == enemies[i].y {
// 			enemies[i] = enemies[len(enemies)-1]
// 			enemies = enemies[:len(enemies)-1]
// 			if len(enemies) == 0 {
// 				enemies = []*enemy{}
// 			}
// 		}
// 	}
// }
