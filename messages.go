package nexmo

import (
	"context"
	"fmt"
	"net/url"
)

const messagesPathPart = "sms/json"

// MessageService is the client holding messaging capatibilites
type MessageService struct {
	client *Client
}

// Message holds message response data
type Message struct {
	MessageCount string `json:"message-count"`
	Messages     []struct {
		To               string `json:"to"`
		MessageID        string `json:"message-id"`
		Status           string `json:"status"`
		ClientReference  string `json:"client-ref"`
		RemainingBalance string `json:"remaining-balance"`
		MessagePrice     string `json:"message-price"`
		Network          string `json:"network"`
		ErrorText        string `json:"error-text"`
	} `json:"messages"`
}

// Create creates a message resource
func (m *MessageService) Create(ctx context.Context, data url.Values) (*Message, error) {
	msg := new(Message)
	err := m.client.CreateResource(ctx, messagesPathPart, data, msg)
	// TODO(lopezator) nexmo returning 200 OK and message response but status != 0 on some errors
	// Suggestions to handle this better are welcome
	// https://developer.nexmo.com/api/sms#errors
	if msg.Messages[0].Status != "0" {
		return nil, fmt.Errorf("an error ocurred %s", msg.Messages[0].ErrorText)
	}

	return msg, err
}

// SendMessage sends an outbound Message with the given text.
func (m *MessageService) SendMessage(from string, to string, text string) (*Message, error) {
	v := url.Values{}
	v.Set("text", text)
	v.Set("from", from)
	v.Set("to", to)

	return m.Create(context.Background(), v)
}
