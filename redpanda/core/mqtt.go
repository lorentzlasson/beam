package core

import (
	"encoding/json"
	"fmt"
	mqtt "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"github.com/lorentzlasson/beam/redpanda/model"
	"github.com/lorentzlasson/beam/redpanda/util/vcapservices"
	"log"
)

var client *mqtt.Client

func startSubscriptions() {
	creds := vcapservices.GetCredentials("iotf-service")

	broker := fmt.Sprintf("tcp://%s:%s", creds["mqtt_host"], creds["mqtt_u_port"])
	clientId := fmt.Sprintf("a:%s:%s%s", creds["org"], "redpanda-", config.appId) // appId to prevent iotf collision

	opts := mqtt.NewClientOptions().AddBroker(broker)
	opts.SetUsername(creds["apiKey"])
	opts.SetPassword(creds["apiToken"])
	opts.SetClientID(clientId)

	log.Printf("Connecting to mqtt broker: %s", broker)

	client = mqtt.NewClient(opts)

	token := client.Connect()
	token.Wait()
	token = client.Subscribe("iot-2/type/app/id/+/evt/new_beam/fmt/json", 0, messageReceived)
	token.Wait()

	// log.Println("Connected to mqtt broker")
}

var messageReceived mqtt.MessageHandler = func(client *mqtt.Client, msg mqtt.Message) {
	log.Printf("TOPIC: %s\n", msg.Topic())
	log.Printf("MSG: %s\n", msg.Payload())
	beacon := &model.Beacon{}
	json.Unmarshal(msg.Payload(), &beacon)
	storeBeacon(beacon)
}
