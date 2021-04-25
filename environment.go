package csl_discovery_client

import env "github.com/carrot-systems/csl-env"

func loadDataFromEnv(appName string) (DiscoveryClientConfiguration, DiscoveryServerConfiguration) {
	clientConfiguration := DiscoveryClientConfiguration{
		Name:        appName,
		ExternalUrl: env.RequireEnvString("EXTERNAL_URL"),
	}

	serverConfiguration := DiscoveryServerConfiguration{
		ServerUrl: env.RequireEnvString("DISCOVERY_SERVER_URL"),
	}

	return clientConfiguration, serverConfiguration
}
