package llm

// func getWeather(city string, street string) string {
// 	if city == "" {
// 		panic("Something went very wrong")
// 	}
// 	fmt.Println("street", street)
// 	if street == "" {
// 		return "34f"
// 	}
// 	return "-3f"
// }
//
// func getWeatherWrapper(args map[string]any) (string, error) {
// 	city, ok := args["city"].(string)
// 	if !ok {
// 		// TODO: Add args to error stdout
// 		fmt.Printf("Error: Unable to get weather data from args")
// 		return "Failed to get weather data", errors.New("Failed to get weather data")
// 	}
// 	street := args["street"].(string)
//
// 	return getWeather(city, street), nil
// }

// output, err := llm.Call([]llm.Message{
// 	llm.NewMessage("Talk like a pirate", llm.SystemMessage),
// 	llm.NewMessage("What is the weather in Chicago?", llm.UserMessage),
// }, []llm.Tool{
// 	llm.NewTool("get_weather", "Get the weather for the user. If street is unknown, do not include in argument", getWeatherWrapper,
// 		llm.NewParameter("city", "string", true),
// 		llm.NewParameter("street", "string", false)),
// })

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/responses"
)

type EmptyStruct struct {}

// func Call(messages []Message) (string, error) {
// 	// Transform the messages to OpenAI messages
// 	openaiMessages := make([]openai.ChatCompletionMessageParamUnion, len(messages))
// 	for i, message := range messages {
// 		var err error
// 		openaiMessages[i], err = message.toOpenAI()
// 		if err != nil {
// 			trimmedMessage := message.message
// 			if len(trimmedMessage) > 15 {
// 				trimmedMessage = trimmedMessage[:15] + "..."
// 			}
// 			return "", fmt.Errorf("There was an issue converting message (%s) from messages to OpenAI format: %v", trimmedMessage, err)
// 		}
// 	}
//
//
// 	client := openai.NewClient()
//
// 	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
// 		Messages: openaiMessages,
// 		Model: openai.ChatModelGPT5_2,
// 	})
//
// 	if err != nil {
// 		return "", fmt.Errorf("Failed to call to OpenAI LLM: %w", err)
// 	}
//
// 	return chatCompletion.Choices[0].Message.Content, nil
// }

func Call(messages []Message, tools []Tool) (string, error) {
	openaiMessages, err := messagesToOpenAIMessages(messages)
	if err != nil {
		return "", err
	}


	toolNamesMap, err := getToolNamesMap(tools)
	if err != nil {
		return "", err
	}

	client := openai.NewClient()

	ctx := context.Background()

	params := responses.ResponseNewParams{
		Model: openai.ChatModelGPT5_2,
		Temperature: openai.Float(0),
		Input: responses.ResponseNewParamsInputUnion{
			OfInputItemList: openaiMessages,
		},
		Tools: toolsToOpenAITools(tools),
	}

	response, _ := client.Responses.New(ctx, params)

	toolResponses := make([]responses.ResponseInputItemUnionParam, 0)
	foundToolRequest := false
	// Check for function calls in the response output
	for _, item := range response.Output {
		if item.Type == "function_call" {
			foundToolRequest = true
			toolCall := item.AsFunctionCall()
			toolName := toolCall.Name

			tool, ok := toolNamesMap[toolName]
			if !ok {
				fmt.Printf("Warning: Non-existent tool call (%s) found in LLM response, skipping tool invocation\n", toolName)
				continue
			}


			var args map[string]any
			json.Unmarshal([]byte(toolCall.Arguments), &args)
			toolResponse, err := tool.call(args)
			if err != nil {
				if toolResponse == "" {
					toolResponse = "Failed to invoke this tool call. Do not try again, answer the best you can"
				}
				fmt.Printf("Error: Tool call (%s) failed, reporting failure in tool response. Error: %v\n", toolName, err)
			}

			toolResponses = append(toolResponses, responses.ResponseInputItemUnionParam{
				OfFunctionCallOutput: &responses.ResponseInputItemFunctionCallOutputParam{
					CallID: toolCall.CallID,
					Output: responses.ResponseInputItemFunctionCallOutputOutputUnionParam{
						OfString: openai.String(toolResponse),
					},
				},
			})
			}
		}

		if foundToolRequest {
			response, err = client.Responses.New(ctx, responses.ResponseNewParams{
				Model:              openai.ChatModelGPT5_2,
				PreviousResponseID: openai.String(response.ID),
				Input: responses.ResponseNewParamsInputUnion{
					OfInputItemList: toolResponses,
				},
			})
			if err != nil {
				return "", fmt.Errorf("There was an issue invoking the post-tool call LLM request: %v", err)
			}
		}

		return response.OutputText(), nil
}

func toolsToOpenAITools(tools []Tool) []responses.ToolUnionParam {
	openaiTools := make([]responses.ToolUnionParam, len(tools))
	for i, tool := range tools {
		openaiTools[i] = tool.toOpenAI()
	}

	return openaiTools
}

func messagesToOpenAIMessages(messages []Message) ([]responses.ResponseInputItemUnionParam, error) {
	// Transform the messages to OpenAI messages
	openaiMessages := make([]responses.ResponseInputItemUnionParam, len(messages))
	for i, message := range messages {
		var err error
		openaiMessages[i], err = message.toOpenAI()
		if err != nil {
			trimmedMessage := message.message
			if len(trimmedMessage) > 15 {
				trimmedMessage = trimmedMessage[:15] + "..."
			}
			return nil, fmt.Errorf("There was an issue converting message (%s) from messages to OpenAI format: %v", trimmedMessage, err)
		}
	}

	return openaiMessages, nil
}

func getToolNamesMap(tools []Tool) (map[string]Tool, error) {
	namesMap := map[string]Tool{}
	for _, tool := range tools {
		if _, exists := namesMap[tool.Name]; exists {
			return map[string]Tool{}, fmt.Errorf("Duplicate name (%s) found in tools", tool.Name)
		}
		namesMap[tool.Name] = tool
	}

	return namesMap, nil
}
