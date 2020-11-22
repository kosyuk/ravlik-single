package main

type RandomRavlik struct {
	Status RavlikStatus
}

var RandomRavlikColony = AbstractColony{
	name: "RandomRavlik#1",
	createRavlik: func() Ravlik {
		return &RandomRavlik{NEW}
	},
}

func (ravlik *RandomRavlik) NextStatus() {
	ravlik.Status = GetRavlikRandomStatus(ravlik.Status)
}
func (ravlik *RandomRavlik) GetStatus() RavlikStatus {
	return ravlik.Status
}
func (ravlik *RandomRavlik) SetStatus(status RavlikStatus) {
	ravlik.Status = status
}
