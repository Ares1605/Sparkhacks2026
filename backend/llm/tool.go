package llm

import (
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/responses"
)

type Parameter struct {
	name string
	type_ string
	isRequired bool
}

func NewParameter(name string, type_ string, isRequired bool) Parameter {
	return Parameter{
		name: name,
		type_: type_,
		isRequired: isRequired,
	}
}

type Tool struct {
	Name string
	Description string
	Parameters []Parameter
	ToolCall func(args map[string]any) (string, error)
}

func NewTool(name string, description string, toolCall func(args map[string]any) (string, error), parameters ...Parameter) Tool {
	return Tool{
		Name: name,
		Description: description,
		Parameters: parameters,
		ToolCall: toolCall,
	}
}

func (t *Tool) call(args map[string]any) (string, error) {
	return t.ToolCall(args)
}

func (t *Tool) toOpenAI() responses.ToolUnionParam {
	properties := map[string]any{}
	required := make([]string, 0)

	for _, parameter := range t.Parameters {
		properties[parameter.name] = map[string]string{
			"type": parameter.type_,
		}

		if parameter.isRequired {
			required = append(required, parameter.name)
		}
	}

	parameters := map[string]any{
		"type": "object",
		"properties": properties,
		"required": required,
	}

	return responses.ToolUnionParam{
			OfFunction: &responses.FunctionToolParam{
				Name: t.Name,
				Description: openai.String(t.Description),
				Parameters: parameters,
			},
		}
}
