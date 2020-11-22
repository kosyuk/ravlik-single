package main

import "fmt"

// Initial Ravlik Count
const (
	RAVLIKCOUNT = 1000000
)

var area Area
var colonyList []Colony

// var ravliks []Ravlik

func init() {
	colonyList = append(colonyList, &RandomRavlikColony, &RandomRavlikColony2, &RavlikWithMemoryColony, &RavlikShareWithKidsColony)
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
	area.status = AreaNextStatus(area.status)
	area.turn++
}

func AreaAffect() {
	for i := range colonyList {
		colonyList[i].AreaAffect(area.status)
	}
}

// DisplayStatistics displays count of ravliks by status
func DisplayStatistics() {
	fmt.Printf("%d.Stats:(%s)\n", area.turn, area.status.String())
	for i := range colonyList {
		fmt.Printf("%s(%d)\n", colonyList[i].GetName(), colonyList[i].GetCount())
		stats := colonyList[i].GetStats()
		for j := range stats {
			fmt.Printf("%s:%d  ", RavlikStatus(j).String(), stats[j])
		}
		fmt.Println()
	}
	fmt.Println()
}
