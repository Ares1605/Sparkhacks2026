package main

import (
	"net/http"
	"x/llm"
)

func recommendations_handler(w http.ResponseWriter, r *http.Request) {
	var startPrompt = `You are Precog, a proactive shopping assistant. Your role is to help the user research, compare, and choose products, find the best deals, suggest alternatives, and warn about stock, prices, compatibility, or scams. Use all available tools and context to anticipate needs, provide reminders, and give concise, actionable advice tailored to the user’s preferences, budget, and trends. All the data must come from one of defined provider, so if you have an idea of what you want to recommend, YOU MUST RETURN REAL RESULTS FROM A PROVIDER, use your search tools. YOUR RESPONSE SHOULD BE A LIST WHERE EACH ITEM FOLLOWS THIS SCHEMA:
		{
			name: string;
			image_path: string;
			provider_url: string;
			price: string;
		} or {
			error: string
		}

DO NOT RESPOND WITH ANY OTHER TEXT THAN THE JSON SCHEMA`

	var userPrompt = `Based on all the data you have about me, provide general recommendations I might benefit from. These can include suggestions, warnings, reminders, or optimizations across any aspect of my shopping habits, preferences, budget, or trends. Highlight what actions or changes could improve my choices or save me time or money.`

	response, err := llm.Call([]llm.Message{
		llm.NewMessage(startPrompt, llm.SystemMessage),
		llm.NewMessage(userPrompt, llm.UserMessage),
	}, build_llm_tools(r.Context()))

	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// type Recommendation struct {
	// 	Name        string `json:"name"`
	// 	ImagePath   string `json:"image_path"`
	// 	ProviderURL string `json:"provider_url"`
	// 	Price       string `json:"price"`
	// }
	// response should be a []Recommendation
	writeResponse(w, http.StatusOK, response)
}
