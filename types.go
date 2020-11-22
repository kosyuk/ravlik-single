package main

// RavlikStatus is enum for statuses of ravlik
type RavlikStatus int

const (
	//NEW is ravlik currently birth
	NEW RavlikStatus = iota
	//IDLE is ravlik is in idle state
	IDLE
	//PARENTING is ravlik in parenting state
	PARENTING
	//DEFENCE is ravlik defencing on area issues
	DEFENCE
	//DEAD is ravlik currenly dead, will be cleaned
	DEAD
)

func (d RavlikStatus) String() string {
	return [...]string{"NEW", "IDLE", "PARENTING", "DEFENCE", "DEAD"}[d]
}

// AreaStatus is enum for statuses for area
type AreaStatus int

// Area status
const (
	SUNNY AreaStatus = iota
	CLOUDY
	STORMY
)

func (d AreaStatus) String() string {
	return [...]string{"SUNNY", "CLOUDY", "STORMY"}[d]
}

type Ravlik interface {
	NextStatus()
	GetStatus() RavlikStatus
	SetStatus(RavlikStatus)
}

type Colony interface {
	Init(int)
	Turn()
	AreaAffect(AreaStatus)
	GetName() string
	GetStats() map[RavlikStatus]int
	GetCount() int
}

// Area Structure
type Area struct {
	turn   int
	status AreaStatus
}
