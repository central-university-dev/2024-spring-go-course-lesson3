package main

import (
	"fmt"
)

func main() {
	player := Player{ID: 1, Name: "Tracer", Team: "blue"}
	player.position.x = 1002
	
	if player.position.isOutOfBounds() {
		player.position = spawn()
	}
	fmt.Println(player)
}

type Player struct {
	ID   int
	Name string
	Team string

	position position
}

type position struct {
	x int
	y int
	z int
}

func (p *position) isOutOfBounds() bool {
	return p.x > 1000 || p.y > 1000 || p.z > 1000
}

func spawn() position {
	return position{x: 21, y: 31, z: 0}
}