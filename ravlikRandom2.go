package main

type RandomRavlik2 struct {
	Status RavlikStatus
}

var RandomRavlikColony2 = AbstractColony{
	name: "RandomRavlik#2",
	createRavlik: func(parent *Ravlik) Ravlik {
		return &RandomRavlik2{NEW}
	},
}

func (ravlik *RandomRavlik2) NextStatus() {
	ravlik.Status = GetRavlikRandomStatus(ravlik.Status)
}
func (ravlik *RandomRavlik2) GetStatus() RavlikStatus {
	return ravlik.Status
}
func (ravlik *RandomRavlik2) SetStatus(status RavlikStatus) {
	ravlik.Status = status
}
