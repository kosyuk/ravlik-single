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
		AreaTurn()
		AreaAffect()
		DisplayStatistics()
	}
	fmt.Println("Ravlik Ended")
}
