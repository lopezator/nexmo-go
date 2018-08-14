package nexmo

import (
	"context"
	"fmt"
	"net/url"
)

const callsPathPart = "/tts/json"

// Call holds call response data
type Call struct {
	CallID    string `json:"call_id"`
	To        string `json:"to"`
	Status    string `json:"status"`
	ErrorText string `json:"error_text"`
}

// CallService is the client holding calling capatibilites
type CallService struct {
	client *Client
}

// Create creates a call resource
func (c *CallService) Create(ctx context.Context, data url.Values) (*Call, error) {
	call := &Call{}
	err := c.client.CreateResource(ctx, callsPathPart, data, call)
	if err != nil {
		return nil, err
	}
	// TODO(lopezator) nexmo returning 200 OK and call response but status != 0 on some errors
	// Suggestions to handle this better are welcome
	if call.Status != "0" {
		return nil, fmt.Errorf("an error occurred %s", call.ErrorText)
	}

	return call, err
}

// MakeTTSCall makes an outbound Call reproducing using TTS with the given text.
func (c *CallService) MakeTTSCall(from string, to string, text string, repeat string) (*Call, error) {
	v := url.Values{}
	v.Set("text", text)
	v.Set("from", from)
	v.Set("to", to)
	v.Set("repeat", repeat)

	return c.Create(context.Background(), v)
}
