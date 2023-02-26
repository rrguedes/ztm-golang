// --Summary:
//
//	Copy your rcv-func solution to this directory and write unit tests.
//
// --Requirements:
// * Write unit tests that ensure:
//   - Health & energy can not go above their maximums
//   - Health & energy can not go below 0
//   - If any of your  tests fail, make the necessary corrections
//     in the copy of your rcv-func solution file.
//
// --Notes:
// * Use `go test -v ./exercise/testing` to run these specific tests
package testing

import "testing"

func TestSetHealth(t *testing.T) {
	player := Player{maxEnergy: 100, maxHealth: 100, name: "CJ"}
	player.setHealth(50)

	if player.health != 50 {
		t.Errorf("Health should be 50, received %v", player.health)
	}
}

func TestSetEnergy(t *testing.T) {
	player := Player{maxEnergy: 100, maxHealth: 100, name: "CJ"}
	player.setEnergy(50)

	if player.energy != 50 {
		t.Errorf("Energy should be 50, received %v", player.energy)
	}
}
