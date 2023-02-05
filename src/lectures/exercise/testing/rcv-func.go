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
	name      PlayerName
	health    Stat
	maxHealth Stat
	energy    Stat
	maxEnergy Stat
}

func NewPlayer(name PlayerName) Player {
	return Player{name: name, health: 100, energy: 100, maxHealth: 500, maxEnergy: 500}
}

func (player *Player) changeHealth(change int) {
	newHealth := player.health + Stat(change)

	if newHealth > player.maxHealth {
		newHealth = player.maxHealth
	}

	if newHealth < 0 {
		newHealth = 0
	}

	if newHealth >= player.health {
		fmt.Printf("Player %v gained %v health points!\n", player.name, newHealth-player.health)
	} else {
		fmt.Printf("Player %v got damaged of %v health points!\n", player.name, player.health-newHealth)
	}
	player.health = newHealth
}

func (player *Player) changeEnergy(change int) {
	newEnergy := player.energy + Stat(change)

	if newEnergy > player.maxEnergy {
		newEnergy = player.maxEnergy
	}

	if newEnergy < 0 {
		newEnergy = 0
	}

	if newEnergy >= player.energy {
		fmt.Printf("Player %v gained %v energy points!\n", player.name, newEnergy-player.energy)
	} else {
		fmt.Printf("Player %v got damaged of %v energy points!\n", player.name, player.energy-newEnergy)
	}
	player.energy = newEnergy
}

func main() {
	readyPlayerOne := NewPlayer("ReadyPlayer1")
	fmt.Println("Initial:", readyPlayerOne)

	readyPlayerOne.changeHealth(10)

	readyPlayerOne.changeHealth(-69)

	readyPlayerOne.changeHealth(99999)

	readyPlayerOne.changeHealth(-99999)

	readyPlayerOne.changeEnergy(69)

	readyPlayerOne.changeEnergy(-1690)

	readyPlayerOne.changeEnergy(-10)
}
