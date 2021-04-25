package csl_discovery_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (client *DiscoveryClient) Register(appName string) error {
	if client.registered {
		return errors.New("client is already registered")
	}

	clientConf, serverConf := loadDataFromEnv(appName)
	client.client, client.server = &clientConf, &serverConf

	dt, err := json.Marshal(&client.client)

	if err != nil {
		return err
	}

	url := fmt.Sprintf("http://%s/services", serverConf.ServerUrl)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(dt))

	if err != nil {
		return err
	}

	decoder := json.NewDecoder(resp.Body)

	var registerResponse RegisterResponse
	err = decoder.Decode(&registerResponse)

	if err != nil {
		return err
	}

	client.registered = registerResponse.Status.Success
	println(registerResponse.Status.Message)
	if !registerResponse.Status.Success {
		return errors.New(registerResponse.Status.Message)
	}
	return nil
}
