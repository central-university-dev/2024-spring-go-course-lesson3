package main

import (
	"fmt"
)

func main() {
	player := Player{ID: 1, Name: "Tracer", Team: "blue"}
	fmt.Println(player)

	player.Team = "red"
	_ = player.Team
}

type Player struct {
	ID   int
	Name string
	Team string
}
