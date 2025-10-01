package main

import (
	"fmt"
)

func main() {
	player, err := New(1, "Tracer", "red")
	fmt.Println(player, err)
}

func New(id int, name string, team string) (Player, error) {
	if team != "blue" && team != "red" && team != "white" {
		return Player{}, fmt.Errorf("unsupported team: %s", team)
	}
	return Player{ID: id, Name: name, Team: team, position: spawn()}, nil
}

func spawn() position {
	return position{x: 21, y: 31, z: 0}
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
