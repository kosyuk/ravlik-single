package main

import "math/rand"

func Filter(arr []Ravlik, cond func(Ravlik) bool) []Ravlik {
	result := []Ravlik{}
	for i := range arr {
		if cond(arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result
}

func GetRavlikParentingChildsCount() int {
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

func isRavlikAffectedByArea(areaStatus AreaStatus, ravlikStatus RavlikStatus) bool {
	rand100 := rand.Intn(100)
	return rand100 <= DeathChance[areaStatus][ravlikStatus]
}
