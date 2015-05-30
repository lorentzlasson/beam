package vcapservices

import (
	"encoding/json"
	"fmt"
	cfenv "github.com/cloudfoundry-community/go-cfenv"
	"os"
)

type vcapServices struct {
	VCAP_SERVICES map[string][]cfenv.Service
}

var servicesLoaded bool
var services = vcapServices{}

func loadCredentials() {
	appEnv, err := cfenv.Current()

	if err != nil { // local
		file, _ := os.Open("VCAP_SERVICES.json")
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&services)
		if err != nil {
			fmt.Println("ERROR:", err)
		}

	} else {
		out, _ := json.Marshal(appEnv)
		fmt.Println("appEnv: ", string(out))
		services.VCAP_SERVICES = appEnv.Services
	}

	servicesLoaded = true
}

func GetCredentials(serviceName string) map[string]string {
	return GetCredentialsOnIndex(serviceName, 0)
}

func GetCredentialsOnIndex(serviceName string, index int) map[string]string {
	if !servicesLoaded {
		loadCredentials()
	}
	return services.VCAP_SERVICES[serviceName][index].Credentials
}
