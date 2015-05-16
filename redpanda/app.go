package main

import (
	randomdata "github.com/Pallinder/go-randomdata"
	"log"
	"time"
)

var startupStart int64
var appId string

func setup() {
	appId = randomdata.SillyName()
	log.Printf("Starting %s", appId)
	startupStart = time.Now().UnixNano()
}

func main() {
	setup()
	openDb()
	startSubscriptions()
	startServer()
}
