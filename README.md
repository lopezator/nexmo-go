# nexmo-go

WIP Nexmo REST API client for Go.

Hugely inspired by: https://github.com/kevinburke/twilio-go

As it's a WIP it only supports message sending by now.

# Usage

```
const apiKey = "my-api-key"
const apiSecret = "my-api-secret"

// Create a client
client := NewClient(apiKey, apiSecret, nil)

// Send a message
msg, err := client.Messages.SendMessage("ME", "+34666666666", "Sent via go-nexmo")
```


