package db

import (
	"encoding/json"
	"time"
)

type Order struct {
	Id         int
	ProviderId string
	Name       string
	Price      float32
	OrderDate  string
}

type MessageRole int

const (
	UserMessage MessageRole = iota
	ServerMessage
)

type ChatMessage struct {
	Role        MessageRole `json:"role"`
	Message     string      `json:"message"`
	SessionUuid string      `json:"session_uuid"`
	MessageDate time.Time   `json:"message_date"`
}
type DBChatMessage struct {
	Id int `json:"id"`
	ChatMessage
}

func NewChatMessage(role MessageRole, message string, sessionUuid string, messageDate time.Time) ChatMessage {
	return ChatMessage{
		Role:        role,
		Message:     message,
		SessionUuid: sessionUuid,
		MessageDate: messageDate,
	}
}

func (c ChatMessage) MarshalJSON() ([]byte, error) {
	type Alias ChatMessage
	return json.Marshal(&struct {
		MessageDate string `json:"message_date"`
		*Alias
	}{
		MessageDate: c.MessageDate.Format("2006-01-02T15:04:05Z"),
		Alias:       (*Alias)(&c),
	})
}

type Provider struct {
	Id       int
	Name     string
	LastSync string
	Username string
	Password string
}

type ProviderStatus struct {
	LastSync *string
	Username *string
}

type ProviderCredentials struct {
	Username string
	Password string
}
