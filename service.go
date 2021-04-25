package csl_discovery_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (client *DiscoveryClient) GetAllServices() ([]Service, error) {
	if !client.registered {
		return nil, errors.New("client is not registered")
	}

	url := fmt.Sprintf("http://%s/services", client.server.ServerUrl)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(resp.Body)

	var discoverResponse DiscoverResponse
	err = decoder.Decode(&discoverResponse)

	if err != nil {
		return nil, err
	}

	if !discoverResponse.Status.Success {
		return nil, errors.New(discoverResponse.Status.Message)
	}
	return discoverResponse.Services, nil
}

func (client *DiscoveryClient) GetService(appName string) ([]Service, error) {
	if !client.registered {
		return nil, errors.New("client is not registered")
	}

	url := fmt.Sprintf("http://%s/services/%s", client.server.ServerUrl, appName)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(resp.Body)

	var discoverResponse DiscoverResponse
	err = decoder.Decode(&discoverResponse)

	if err != nil {
		return nil, err
	}

	if !discoverResponse.Status.Success {
		return nil, errors.New(discoverResponse.Status.Message)
	}
	return discoverResponse.Services, nil
}
