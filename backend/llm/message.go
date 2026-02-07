package llm

import (
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/responses"
)

type MessageType int

// SystemMessage can be used for system instructions and
// tool calls definition, UserMessage can be used for
// passing the user prompt
const (
	SystemMessage MessageType = iota
	UserMessage
)

type Message struct {
	messageType MessageType
	message string
}

func NewMessage(message string, messageType MessageType) Message {
	return Message{
		message: message,
		messageType: messageType,
	}
}

func (m *Message) toOpenAI() (responses.ResponseInputItemUnionParam, error) {
	switch m.messageType {
	case SystemMessage:
		return responses.ResponseInputItemUnionParam{
					OfMessage: &responses.EasyInputMessageParam{
						Content: responses.EasyInputMessageContentUnionParam{
							OfString: openai.String(m.message),
						},
						Role: responses.EasyInputMessageRoleSystem,
						Type: "message",
					},
				}, nil
	case UserMessage:
		return responses.ResponseInputItemUnionParam{
					OfMessage: &responses.EasyInputMessageParam{
						Content: responses.EasyInputMessageContentUnionParam{
							OfString: openai.String(m.message),
						},
						Role: responses.EasyInputMessageRoleUser,
						Type: "message",
					},
				}, nil
	default:
		return responses.ResponseInputItemUnionParam{}, fmt.Errorf("Message.toOpenAI does not support MessageType (%v)", m.messageType)
	}
}
