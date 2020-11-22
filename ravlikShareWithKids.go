package main

import "reflect"

type RavlikShareWithKids struct {
	Status    RavlikStatus
	Decisions map[RavlikStatus]RavlikStatus
}

var RavlikShareWithKidsColony = AbstractColony{
	name: "ravlikShareWithKids",
	createRavlik: func(parent *Ravlik) Ravlik {
		if parent == nil {
			return &RavlikWithMemory{NEW, make(map[RavlikStatus]RavlikStatus)}
		}
		val := reflect.ValueOf(*parent).Elem()
		parentDecisions := val.FieldByName("Decisions").Interface().(map[RavlikStatus]RavlikStatus)
		return &RavlikWithMemory{NEW, parentDecisions}
	},
}

func (ravlik *RavlikShareWithKids) NextStatus() {
	status, ok := ravlik.Decisions[ravlik.Status]
	if ok {
		ravlik.Status = status
		return
	}
	newStatus := GetRavlikRandomStatus(ravlik.Status)
	ravlik.Decisions[ravlik.Status] = newStatus
	ravlik.Status = newStatus
}

func (ravlik *RavlikShareWithKids) GetStatus() RavlikStatus {
	return ravlik.Status
}

func (ravlik *RavlikShareWithKids) SetStatus(status RavlikStatus) {
	ravlik.Status = status
}
