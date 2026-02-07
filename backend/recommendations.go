package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"x/llm"
)

func recommendations_handler(w http.ResponseWriter, r *http.Request) {

	orders, err := database.GetAllOrder(r.Context())
	if err != nil {
		panic(":(")
	}

	buf, err := json.Marshal(&orders)
	ordersStringified := string(buf)
	var startPrompt = fmt.Sprintf(`You are Precog, a proactive shopping assistant. Your role is to help the user research, compare, and choose products, find the best deals, suggest alternatives, and warn about stock, prices, compatibility, or scams. Use all available tools and context to anticipate needs, provide reminders, and give concise, actionable advice tailored to the user’s preferences, budget, and trends. All the data must come from one of defined providers, so if you have an idea of what you want to recommend, YOU MUST RETURN REAL RESULTS FROM A PROVIDER.

	Here is some user order history from the user. You should use this data to creatively give recommendations.
	Order History: %s

	Here is an example of how you can use order history to give useful recommendations:
		* If the order history shows many orders with household items, perform an amazon search query for a household item they haven't yet bought (ie. A night light), and provide the Amazon search results as structured JSON data from the search_amazon tool call.

	Use the tools named search_amazon to search for Amazon products to give Amazon recommendations. YOUR RESPONSE SHOULD BE A LIST WHERE EACH ITEM FOLLOWS THIS SCHEMA:
		{
			name: string;
			image_path: string;
			provider_url: string;
			price: string;
		} or {
			error: string
		}

DO NOT RESPOND WITH ANY OTHER TEXT THAN THE JSON SCHEMA. Output only a JSON array.`, ordersStringified)

	var userPrompt = `Based on all the data you have about me, provide interesting recommendations I might benefit from. Be ambigious and creative. Do not tell me you cannot answer the question, you have the context and capabilities. These can include suggestions, warnings, reminders, or optimizations across any aspect of my shopping habits, preferences, budget, or trends. Highlight what actions or changes could improve my choices or save me time or money.`

	response, err := llm.Call([]llm.Message{
		llm.NewMessage(startPrompt, llm.SystemMessage),
		llm.NewMessage(userPrompt, llm.UserMessage),
	}, build_llm_tools(r.Context()))

	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	type Recommendation struct {
		Name        string `json:"name"`
		ImagePath   string `json:"image_path"`
		ProviderURL string `json:"provider_url"`
		Price       string `json:"price"`
		Error       string `json:"error"`
	}

	var recommendations []Recommendation
	if err := json.Unmarshal([]byte(response), &recommendations); err != nil {
		log.Printf("Recommendations LLM returned invalid JSON: %v", err)
		writeError(w, http.StatusInternalServerError, "LLM returned invalid JSON")
		return
	}

	writeResponse(w, http.StatusOK, recommendations)
}
