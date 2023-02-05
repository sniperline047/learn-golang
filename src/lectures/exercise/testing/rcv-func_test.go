//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import "testing"

func TestChangeHealth(t *testing.T) {
	player := NewPlayer("ReadyPlayer1")

	player.changeHealth(int(player.maxHealth) + 1)
	if player.health > player.maxHealth {
		t.Errorf("Health of %v cannot be more than %v, Current value: %v\n", player.name, player.maxHealth, player.health)
	}

	player.changeHealth(-(int(player.maxHealth) + 1))
	if player.health < 0 {
		t.Errorf("Health of %v cannot be less than 0, Current value: %v\n", player.name, player.health)
	}
}

func TestChangeEnergy(t *testing.T) {
	player := NewPlayer("ReadyPlayer1")

	player.changeEnergy(int(player.maxHealth) + 1)
	if player.energy > player.maxEnergy {
		t.Errorf("Energy of %v cannot be more than %v, Current value: %v\n", player.name, player.maxEnergy, player.energy)
	}

	player.changeEnergy(-(int(player.maxHealth) + 1))
	if player.energy < 0 {
		t.Errorf("Energy of %v cannot be less than 0, Current value: %v\n", player.name, player.energy)
	}
}
