package core

import (
	"fmt"
	"github.com/lorentzlasson/beam/redpanda/util"
	"log"
	"net/http"
	"os"
	"time"
)

func getHostAndPort() (host string, port string) {
	if port = os.Getenv("VCAP_APP_PORT"); len(port) == 0 {
		port = "8080"
	}

	location := "Bluemix"
	if host = os.Getenv("VCAP_APP_HOST"); len(host) == 0 { // on Bluemix
		location = "local network"
		var err error
		if host, err = util.GetLocalIp(); err != nil { // on local network
			host = "localhost" // no network
			location = host
		}
	}
	log.Printf("Running on %s", location)
	return
}

func handleDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World %s!\n", config.appId)

	fmt.Fprintln(w, "\n<><> Users <><>")
	for _, user := range users() {
		fmt.Fprintf(w, "Id: %d\n", user.Id)
	}

	fmt.Fprintln(w, "\n<><> Beacons <><>")
	for _, beacon := range beacons() {
		fmt.Fprintf(w, "Id: %d\n", beacon.Id)
	}
}

func startServer() {
	host, port := getHostAndPort()
	http.HandleFunc("/", handleDefault)
	log.Printf("Starting app on %+v:%+v\n", host, port)
	startupTimeMillis := (time.Now().UnixNano() - config.startupStart) / int64(time.Millisecond)
	log.Printf("Startup time in millis: %d", startupTimeMillis)
	http.ListenAndServe(host+":"+port, nil)
}
