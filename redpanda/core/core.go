package core

import (
	randomdata "github.com/Pallinder/go-randomdata"
	"log"
	"time"
)

var config struct {
	startupStart int64
	appId        string
}

func setup() {
	config.appId = randomdata.SillyName()
	config.startupStart = time.Now().UnixNano()
	log.Printf("Starting %s", config.appId)
}

func Start() {
	setup()
	openDb()
	startSubscriptions()
	startServer()
}
