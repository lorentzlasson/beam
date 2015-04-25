package bluemix

import (
	"encoding/json"
	"fmt"
	"os"
)

type vcapServices struct {
	Services map[string][]service
}

type service struct {
	Name        string
	Label       string
	Plan        string
	Credentials map[string]interface{}
}

var services = vcapServices{}

func loadCredentials() {
	file, _ := os.Open("VCAP_SERVICES.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&services)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func GetCredentials(serviceName string) map[string]interface{} {
	return GetCredentialsOnIndex(serviceName, 0)
}

func GetCredentialsOnIndex(serviceName string, index int) map[string]interface{} {
	if services.Services == nil {
		loadCredentials()
		fmt.Println("load cred")
	}
	return services.Services[serviceName][index].Credentials
}
