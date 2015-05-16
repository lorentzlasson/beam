package main

import (
	"fmt"
	"github.com/eaigner/jet"
	"github.com/lib/pq"
	"github.com/lorentzlasson/beam/redpanda/util/vcapservices"
	"log"
)

var db *jet.Db

func openDb() {
	credentials := vcapservices.GetCredentials("elephantsql")

	var err error

	pgUrl, err := pq.ParseURL(credentials["uri"])
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connecting to db: %s", "ElephantSQL")

	db, err = jet.Open("postgres", pgUrl)
	if err != nil {
		log.Fatal(err)
	}
	// log.Println("Connected to db")
}

func users() (users []User) {
	db.Query("SELECT * FROM \"user\"").Rows(&users)
	return
}

func beacons() (beacons []Beacon) {
	db.Query("SELECT * FROM \"beacon\"").Rows(&beacons)
	return
}

func storeBeacon(beacon *Beacon) {
	query := fmt.Sprintf("INSERT INTO \"beacon\" (\"id\", \"userId\") values (%d, %d)", beacon.Id, beacon.UserId)
	log.Println(query)
	err := db.Query(query).Run()
	if err != nil {
		log.Println("err", err)
	}
}
