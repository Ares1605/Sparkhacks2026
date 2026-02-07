package main

type amazonOrderSchema struct {
	OrderNum           string `json:"order_number"`
	OrderData          string `json:"order_placed_date"`
	ProviderDetailsURL string `json:"order_details_link"`
	Items              []struct {
		ImageURL string `json:"image_link"`
		Price    int    `json:"price"`
		Quantity int    `json:"quantity"`
		Title    string `json:"title"`
		ItemUrl  string `json:"link"`
	}
}
