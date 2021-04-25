package csl_discovery_client

type Status struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type RegisterResponse struct {
	Status Status `json:"status"`
}

type Service struct {
	Name        string `json:"name"`
	ExternalUrl string `json:"external_url"`
}

type DiscoverResponse struct {
	Status Status `json:"status"`

	Services []Service `json:"services",omitempty`
}
