//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Stat int
type PlayerName string

type Player struct {
	name   PlayerName
	health Stat
	energy Stat
}

func NewPlayer(name PlayerName) Player {
	return Player{name: name, health: 100, energy: 100}
}

func (player *Player) changeHealth(change int) {
	newHealth := player.health + Stat(change)

	if newHealth > 100 {
		newHealth = 100
	}

	if newHealth < 0 {
		newHealth = 0
	}

	player.health = newHealth
}

func (player *Player) changeEnergy(change int) {
	newEnergy := player.energy + Stat(change)

	if newEnergy > 100 {
		newEnergy = 100
	}

	if newEnergy < 0 {
		newEnergy = 0
	}

	player.energy = newEnergy
}

func main() {
	readyPlayerOne := NewPlayer("Ready Player One")
	fmt.Println("Initial:", readyPlayerOne)

	readyPlayerOne.changeHealth(10)
	fmt.Println("Increased Health:", readyPlayerOne)

	readyPlayerOne.changeHealth(-69)
	fmt.Println("Decreased Health:", readyPlayerOne)

	readyPlayerOne.changeEnergy(69)
	fmt.Println("Increased Energy:", readyPlayerOne)

	readyPlayerOne.changeEnergy(-169)
	fmt.Println("Decreased Energy:", readyPlayerOne)
}
