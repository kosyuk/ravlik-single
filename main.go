package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Ravlik Started")

	for i := 0; i < 100; i++ {
		ColoniesTurn()
		// RavlikListTurn()
		AreaTurn()
		AreaAffect()
		// CleanDeadRavlik()
		DisplayStatistics()
	}
	// fmt.Println(AreaOptions)
	// fmt.Println(RavlikOptions)
	// fmt.Println(DeathChance)
	fmt.Println("Ravlik Ended")
}
