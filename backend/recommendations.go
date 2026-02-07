package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"x/llm"
)

func recommendations_handler(w http.ResponseWriter, r *http.Request) {
	var startPrompt = `You are Precog, a proactive shopping assistant. Your role is to help the user research, compare, and choose products, find the best deals, suggest alternatives, and warn about stock, prices, compatibility, or scams. Use all available tools and context to anticipate needs, provide reminders, and give concise, actionable advice tailored to the user’s preferences, budget, and trends.`

	var userPrompt = `Based on all the data you have about me, provide general recommendations I might benefit from. These can include suggestions, warnings, reminders, or optimizations across any aspect of my shopping habits, preferences, budget, or trends. Highlight what actions or changes could improve my choices or save me time or money.`

	type Recommendation struct {
		Name        string `json:"name"`
		ImagePath   string `json:"image_path"`
		ProviderURL string `json:"provider_url"`
		Price       string `json:"price"`
	}

	buf, _ := os.ReadFile("scripts/amazon-mock-data.json")
	var res []amazonOrderSchema
	json.Unmarshal(buf, &res)

	var recommendations []Recommendation
	for _, order := range res {
		for _, order_item := range order.Items {
			recommendations = append(recommendations, Recommendation{
				Name:        order_item.Title,
				ImagePath:   fmt.Sprintf("/resource/image/%s", url.PathEscape(order_item.ImageURL)),
				Price:       strconv.Itoa(order_item.Price),
				ProviderURL: order_item.ItemUrl,
			})
		}
	}

	writeResponse(w, http.StatusOK, recommendations)
	return

	response, err := llm.Call([]llm.Message{
		llm.NewMessage(startPrompt, llm.SystemMessage),
		llm.NewMessage(userPrompt, llm.UserMessage),
	}, build_llm_tools(r.Context()))

	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, response)
}
