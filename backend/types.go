package main

type OkResponse struct {
	Data any `json:"data"`
}

type ErrorResponse struct {
	Err string `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type AmazonProviderDetails struct {
	LoggedIn   bool    `json:"logged_in"`
	LastSynced *string `json:"last_synced"`
	Username   *string `json:"username"`
}

type ProviderDetailsResponse struct {
	Amazon AmazonProviderDetails `json:"amazon"`
}
