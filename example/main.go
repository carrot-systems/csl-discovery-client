package main

import (
	discoveryClient "github.com/carrot-systems/csl-discovery-client"
	env "github.com/carrot-systems/csl-env"
)

func main() {
	env.LoadEnv()
	discovery := discoveryClient.NewClient()
	err := discovery.Register("example")

	if err != nil {
		println("error registering: " + err.Error())
		return
	}

	devices, err := discovery.GetAllServices()

	if err != nil {
		println("error fetching all devices " + err.Error())
		return
	}

	for _, device := range devices {
		println(device.Name + " = " + device.ExternalUrl)
	}

}
