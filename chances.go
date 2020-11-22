package main

import "fmt"

//status count for arrays
const (
	RAVLIKSTATUSCOUNT = 5
	AREASTATUSCOUNT   = 3
	MAXKIDSCOUNT      = 5
)

// AreaOptions describes How Area trunsitions
var AreaOptions [AREASTATUSCOUNT][AREASTATUSCOUNT]int

// RavlikOptions now ravlik transits by default
var RavlikOptions [RAVLIKSTATUSCOUNT][RAVLIKSTATUSCOUNT]int

// DeathChance fill how area affect ravlic depends on status
var DeathChance [AREASTATUSCOUNT][RAVLIKSTATUSCOUNT]int

// ParentingChance fill how area affect ravlic depends on status
var ParentingChance [MAXKIDSCOUNT]int

func init() {
	fillAreaOptions()
	fillRavlikOptions()
	fillDeathChance()
	fillParentingChance()
	fmt.Println("Ravlik Initialised")
}

func fillAreaOptions() {
	AreaOptions[SUNNY][SUNNY] = 80
	AreaOptions[SUNNY][CLOUDY] = 15
	AreaOptions[SUNNY][STORMY] = 5

	AreaOptions[CLOUDY][SUNNY] = 50
	AreaOptions[CLOUDY][CLOUDY] = 15
	AreaOptions[CLOUDY][STORMY] = 45

	AreaOptions[STORMY][SUNNY] = 80
	AreaOptions[STORMY][CLOUDY] = 15
	AreaOptions[STORMY][STORMY] = 5
}

func fillRavlikOptions() {
	RavlikOptions[NEW][NEW] = 0
	RavlikOptions[NEW][IDLE] = 50
	RavlikOptions[NEW][PARENTING] = 0
	RavlikOptions[NEW][DEFENCE] = 30
	RavlikOptions[NEW][DEAD] = 20

	RavlikOptions[IDLE][NEW] = 0
	RavlikOptions[IDLE][IDLE] = 50
	RavlikOptions[IDLE][PARENTING] = 10
	RavlikOptions[IDLE][DEFENCE] = 20
	RavlikOptions[IDLE][DEAD] = 20

	RavlikOptions[PARENTING][NEW] = 0
	RavlikOptions[PARENTING][IDLE] = 50
	RavlikOptions[PARENTING][PARENTING] = 0
	RavlikOptions[PARENTING][DEFENCE] = 30
	RavlikOptions[PARENTING][DEAD] = 20

	RavlikOptions[DEFENCE][NEW] = 0
	RavlikOptions[DEFENCE][IDLE] = 50
	RavlikOptions[DEFENCE][PARENTING] = 20
	RavlikOptions[DEFENCE][DEFENCE] = 10
	RavlikOptions[DEFENCE][DEAD] = 20

	RavlikOptions[DEAD][NEW] = 0
	RavlikOptions[DEAD][IDLE] = 0
	RavlikOptions[DEAD][PARENTING] = 0
	RavlikOptions[DEAD][DEFENCE] = 0
	RavlikOptions[DEAD][DEAD] = 100
}

func fillDeathChance() {
	DeathChance[SUNNY][NEW] = 10
	DeathChance[SUNNY][IDLE] = 10
	DeathChance[SUNNY][PARENTING] = 20
	DeathChance[SUNNY][DEFENCE] = 20
	DeathChance[SUNNY][DEAD] = 0

	DeathChance[CLOUDY][NEW] = 10
	DeathChance[CLOUDY][IDLE] = 10
	DeathChance[CLOUDY][PARENTING] = 20
	DeathChance[CLOUDY][DEFENCE] = 20
	DeathChance[CLOUDY][DEAD] = 0

	DeathChance[STORMY][NEW] = 50
	DeathChance[STORMY][IDLE] = 50
	DeathChance[STORMY][PARENTING] = 60
	DeathChance[STORMY][DEFENCE] = 10
	DeathChance[STORMY][DEAD] = 0
}

// Probability of ravlik Count that can be born at the same time
func fillParentingChance() {
	ParentingChance[0] = 60
	ParentingChance[1] = 20
	ParentingChance[2] = 10
	ParentingChance[3] = 7
	ParentingChance[4] = 3
}
