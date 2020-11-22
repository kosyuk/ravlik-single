package main

import "math/rand"

type RandomRavlik struct {
	status RavlikStatus
}
type TRandomRavlikColony struct {
	ravlikList []Ravlic
	count      int
}

var RandomRavlikColony TRandomRavlikColony

func (ravlik *RandomRavlik) nextStatus() {
	currentStatus := ravlik.status
	rand100 := rand.Intn(100)
	sum := 0
	for i := 0; i < len(RavlikOptions[currentStatus]); i++ {
		sum += RavlikOptions[currentStatus][i]
		if sum >= rand100 {
			ravlik.status = RavlikStatus(i)
			return
		}
	}
	ravlik.status = currentStatus
}

func (colony *TRandomRavlikColony) Init(count int) {
	for i := 0; i < count; i++ {
		ravlik := RandomRavlik{NEW}
		colony.ravlikList = append(colony.ravlikList, &ravlik)
	}
}

func (colony *TRandomRavlikColony) Turn() {
	for i := range colony.ravlikList {
		colony.ravlikList[i].nextStatus()
	}
}
