package main

type AbstractColony struct {
	ravlikList   []Ravlik
	count        int
	name         string
	createRavlik func() Ravlik
}

func (colony *AbstractColony) GetName() string {
	return colony.name
}

func (colony *AbstractColony) Init(count int) {
	for i := 0; i < count; i++ {
		colony.ravlikList = append(colony.ravlikList, colony.createRavlik())
	}
}

func (colony *AbstractColony) Turn() {
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
				newRavlikList = append(newRavlikList, colony.createRavlik())
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

func (colony *AbstractColony) AreaAffect(areaStatus AreaStatus) {
	for i := range colony.ravlikList {
		if isRavlikAffectedByArea(areaStatus, colony.ravlikList[i].GetStatus()) {
			colony.ravlikList[i].SetStatus(DEAD)
		}
	}
}

func (colony *AbstractColony) GetStats() map[RavlikStatus]int {
	var statusMap = make(map[RavlikStatus]int)
	for i := range colony.ravlikList {
		statusMap[colony.ravlikList[i].GetStatus()]++
	}
	return statusMap
}

func (colony *AbstractColony) GetCount() int {
	return colony.count
}
