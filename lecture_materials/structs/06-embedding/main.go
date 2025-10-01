package main

import (
	"fmt"
)

func main() {
	
	player := Player{Entity: Entity{ID: 1, position: position{1001, 310, 450}}, Name: "Tracer", Team: "blue"}
	if player.isOutOfBounds() {
		player.reset()
	}


	fmt.Println(player)
}

type Entity struct {
	ID       int
	position position
}

func (e *Entity) isOutOfBounds() bool {
	return e.position.x > 1000 || e.position.y > 1000 || e.position.z > 1000
}

func (e *Entity) reset() {
	e.position = position{0, 0, 0}
}

type Player struct {
	Entity

	Name string
	Team string
}

func (e *Player) reset() {
	e.position = spawn()
}

type position struct {
	x int
	y int
	z int
}

func spawn() position {
	return position{x: 21, y: 31, z: 0}
}
