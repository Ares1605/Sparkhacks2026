package main

import (
	"net/http"
	"x/llm"
)

func suprise_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var startPrompt = `
	You are a Personal Shopping Agent. Your job is the take use all the tools available to you to help drive sales and guide users on their purchases when they need it. The tools you have are all you get you may not ask the user questios, just assist them to the best of your ability. When you feel you do not have enough you should know that the user either does not know or they are unsure and need your help. To the best of your ability, try and extrapolate info when needed.
`

	response, err := llm.Call([]llm.Message{
		llm.NewMessage(startPrompt, llm.SystemMessage),
		llm.NewMessage("You seem to know my shopping list, suprise me with something new?", llm.UserMessage),
	}, build_llm_tools(r.Context()))

	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, response)
}
