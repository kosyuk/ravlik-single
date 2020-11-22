package main

type RavlikWithMemory struct {
	Status    RavlikStatus
	Decisions map[RavlikStatus]RavlikStatus
}

var RavlikWithMemoryColony = AbstractColony{
	name: "RavlikWithMemory",
	createRavlik: func(parent *Ravlik) Ravlik {
		return &RavlikWithMemory{NEW, make(map[RavlikStatus]RavlikStatus)}
	},
}

func (ravlik *RavlikWithMemory) NextStatus() {
	status, ok := ravlik.Decisions[ravlik.Status]
	if ok {
		ravlik.Status = status
		return
	}
	newStatus := GetRavlikRandomStatus(ravlik.Status)
	ravlik.Decisions[ravlik.Status] = newStatus
	ravlik.Status = newStatus
}

func (ravlik *RavlikWithMemory) GetStatus() RavlikStatus {
	return ravlik.Status
}

func (ravlik *RavlikWithMemory) SetStatus(status RavlikStatus) {
	ravlik.Status = status
}
