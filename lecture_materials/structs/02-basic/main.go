package main

import (
	"fmt"
)

func main() {
	player := Player{ID: 1, Name: "Tracer", Team: "blue"}
	
	fmt.Println(player)

	player.position = position{x: 1, y: 3, z: -3}
	fmt.Println(player)
}

type Player struct {
	ID   int
	Name string
	Team string // с прописной буквы = публичные

	position position // со строчной буквы = приватные
}

type position struct {
	x int
	y int
	z int
}
