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

package testing

import "fmt"

type Player struct {
	health, maxHealth, energy, maxEnergy int
	name                                 string
}

func (player *Player) setHealth(health int) {
	player.health = health
}

func (player *Player) setEnergy(energy int) {
	player.energy = energy
}

func main() {
	player := Player{maxHealth: 100, maxEnergy: 100, name: "CJ"}
	fmt.Println("Player before functions", player)
	player.setEnergy(50)
	player.setHealth(70)
	fmt.Println("Player after functions", player)
}
