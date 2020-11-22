package main

import (
	"fmt"
	"math/rand"
)

// Initial Ravlik Count
const (
	RAVLIKCOUNT = 1000
)

var area Area
var colonyList []Colony

// var ravliks []Ravlik

func init() {
	colonyList = append(colonyList, &RandomRavlikColony)
	area.status = SUNNY
	for i := range colonyList {
		colonyList[i].Init(RAVLIKCOUNT)
	}
}

func ColoniesTurn() {
	for i := range colonyList {
		colonyList[i].Turn()
	}
}

// AreaTurn provides trn for area
func AreaTurn() {
	area.status = areaNextStatus(area.status)
}

func ravlikNextRandomStatus(currentStatus RavlikStatus) RavlikStatus {
	rand100 := rand.Intn(100)
	sum := 0
	for i := 0; i < len(RavlikOptions[currentStatus]); i++ {
		sum += RavlikOptions[currentStatus][i]
		if sum >= rand100 {
			return RavlikStatus(i)
		}
	}
	return currentStatus
}

func getRavlikParentingChildsCount() int {
	rand100 := rand.Intn(100)
	sum := 0
	for i := 0; i < len(ParentingChance); i++ {
		sum += ParentingChance[i]
		if sum >= rand100 {
			return i
		}
	}
	return 0
}

func ravlikParenting() {
	c := getRavlikParentingChildsCount()
	for i := 0; i < c; i++ {
		ravlik := Ravlik{int64(i), NEW}
		ravliks = append(ravliks, ravlik)
	}
}

// RavlikTurn change status of ravlik
func RavlikTurn(r *Ravlik) {
	r.status = ravlikNextRandomStatus(r.status)
}

// RavlikListTurn provides turn for ravlik lis
func RavlikListTurn() {
	for i := 0; i < len(ravliks); i++ {
		RavlikTurn(&ravliks[i])
	}
}

// CleanDeadRavlik cleans ravlik list from dead ravliks
func CleanDeadRavlik() {

}

// DisplayStatistics displays count of ravliks by status
func DisplayStatistics() {
	fmt.Println("Statistyc of ravliks: ", len(ravliks))
	var statusMap [RAVLIKSTATUSCOUNT]int
	for i := 0; i < len(ravliks); i++ {
		statusMap[ravliks[i].status]++
	}

	fmt.Println(statusMap)
}
