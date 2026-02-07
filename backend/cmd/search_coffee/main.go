package main

import (
	"context"
	"encoding/json"
	"os"

	"x/search"
)

func main() {
	service, err := search.NewServiceFromEnv()
	if err != nil {
		panic(err)
	}

	items, err := service.SearchAmazon(context.Background(), "coffee beans", 5)
	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(items); err != nil {
		panic(err)
	}
}
