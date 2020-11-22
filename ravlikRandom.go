package main

import "math/rand"

type RandomRavlik struct {
	Status RavlikStatus
}
type TRandomRavlikColony struct {
	ravlikList []Ravlik
	count      int
}

var RandomRavlikColony TRandomRavlikColony

func (ravlik *RandomRavlik) NextStatus() {
	currentStatus := ravlik.Status
	rand100 := rand.Intn(100)
	sum := 0
	for i := 0; i < len(RavlikOptions[currentStatus]); i++ {
		sum += RavlikOptions[currentStatus][i]
		if sum >= rand100 {
			ravlik.Status = RavlikStatus(i)
			return
		}
	}
	ravlik.Status = currentStatus
}
func (ravlik *RandomRavlik) GetStatus() RavlikStatus {
	return ravlik.Status
}
func (ravlik *RandomRavlik) SetStatus(status RavlikStatus) {
	ravlik.Status = status
}

func (colony *TRandomRavlikColony) GetName() string {
	return "Random Colony"
}

func (colony *TRandomRavlikColony) Init(count int) {
	for i := 0; i < count; i++ {
		colony.ravlikList = append(colony.ravlikList, &RandomRavlik{NEW})
	}
}

func (colony *TRandomRavlikColony) Turn() {
	//remove dead ravliks
	colony.ravlikList = Filter(colony.ravlikList, func(val Ravlik) bool {
		return val.GetStatus() != DEAD
	})
	//parenting ravliks
	var newRavlikList []Ravlik
	for i := range colony.ravlikList {
		if colony.ravlikList[i].GetStatus() == PARENTING {
			count := GetRavlikParentingChildsCount()
			for j := 0; j < count; j++ {
				newRavlikList = append(newRavlikList, &RandomRavlik{NEW})
			}
		}
	}
	//change status ravliks
	for i := range colony.ravlikList {
		colony.ravlikList[i].NextStatus()
	}
	//add parenting ravliks to list
	colony.ravlikList = append(colony.ravlikList, newRavlikList...)
	colony.count = len(colony.ravlikList)
}

func (colony *TRandomRavlikColony) AreaAffect(areaStatus AreaStatus) {
	for i := range colony.ravlikList {
		if isRavlikAffectedByArea(areaStatus, colony.ravlikList[i].GetStatus()) {
			colony.ravlikList[i].SetStatus(DEAD)
		}
	}
}

func (colony *TRandomRavlikColony) GetStats() map[RavlikStatus]int {
	var statusMap = make(map[RavlikStatus]int)
	for i := range colony.ravlikList {
		statusMap[colony.ravlikList[i].GetStatus()]++
	}
	return statusMap
}
func (colony *TRandomRavlikColony) GetCount() int {
	return colony.count
}
