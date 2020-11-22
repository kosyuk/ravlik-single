package main

import (
	"math/rand"
)

func AreaNextStatus(currentStatus AreaStatus) AreaStatus {
	rand100 := rand.Intn(100)
	sum := 0
	for i := 0; i < len(AreaOptions[currentStatus]); i++ {
		sum += AreaOptions[currentStatus][i]
		if sum >= rand100 {
			return AreaStatus(i)
		}
	}
	return currentStatus
}
