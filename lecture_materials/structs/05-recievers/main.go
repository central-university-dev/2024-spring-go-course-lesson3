package main

import (
	"fmt"
)

func main() {
	player := Player{Name: "Ana", Team: "white"}

	player.SetName("Tracer")
	player.SetTeam("blue")
	fmt.Println(player.GetName(), player.GetTeam()) // Tracer, white

	// access via pointer
	ptr := &player
	ptr.SetName("Player")
	ptr.SetTeam("red")
	fmt.Println(ptr.GetName(), ptr.GetTeam()) // Player, white
}

type Player struct {
	ID   int
	Name string
	Team string
}

func (p *Player) GetName() string {
	return p.Name
}

func (p *Player) SetName(name string) {
	p.Name = name
}

func (p Player) GetTeam() string {
	return p.Team
}

func (p Player) SetTeam(team string) {
	p.Team = team
}
