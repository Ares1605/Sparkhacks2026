package main

type OkResponse struct {
	Data any `json:"data"`
}

type ErrorResponse struct {
	Err string `json:"error"`
}

type ProviderDetails struct {
	Name       string `json:"name"`
	LastSynced string `json:"last_synced"`
}
