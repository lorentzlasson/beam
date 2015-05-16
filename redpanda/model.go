package main

import (
	"time"
)

type User struct {
	Id int
}

type Beacon struct {
	Id     int
	Lat    float64
	Lng    float64
	UserId int
	Time   *time.Time
}
