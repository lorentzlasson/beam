package model

import (
	"time"
)

type Beacon struct {
	Id     int
	Lat    float64
	Lng    float64
	UserId int
	Time   *time.Time
}

func (b *Beacon) add() {

}
