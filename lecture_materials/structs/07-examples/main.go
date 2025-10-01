package main

import (
	"fmt"
)

func main() {
	wall := Entity{ID: 2, position: struct {
		x int
		y int
		z int
	}{1, 3, 4}}
	if wall.isOutOfBounds() {
		wall.reset()
	}
	fmt.Println(wall)

	lake := Entity{ID: 2, position: struct{x int; y int; z int}{2, 4, 6}}
	lake2 := Entity{ID: 2, position: struct{x int; y int; z int}{2, 4, 6}}
	fmt.Println(lake == wall)
	fmt.Println(lake == lake2)
	fmt.Printf("lake: %p; lake2: %p\n", &lake, &lake2)
}

type Entity struct {
	ID int

	position struct {
		x int
		y int
		z int
	}

	// Parent *Entity
	// Siblings []Entity
}

func (e *Entity) isOutOfBounds() bool {
	return e.position.x > 1000 || e.position.y > 1000 || e.position.z > 1000
}

func (e *Entity) reset() {
	e.position = struct {
		x int
		y int
		z int
	}{0, 0, 0}
}
