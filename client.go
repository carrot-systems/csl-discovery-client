package csl_discovery_client

type DiscoveryClientConfiguration struct {
	Name        string `json:"name"`
	ExternalUrl string `json:"external_url"`
}

type DiscoveryClient struct {
	RegisteredName string

	registered bool
	server     *DiscoveryServerConfiguration
	client     *DiscoveryClientConfiguration
}

func NewClient() *DiscoveryClient {
	return &DiscoveryClient{
		registered: false,
		server:     nil,
		client:     nil,
	}
}
