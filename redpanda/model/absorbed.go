package model

import (
	"time"
)

type Absorbed struct {
	Id       int
	Time     *time.Time
	BeaconId int
	BeamId   int
}
