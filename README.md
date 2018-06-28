[![Build Status](https://travis-ci.org/lopezator/nexmo-go.svg?branch=master)](https://travis-ci.org/lopezator/nexmo-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/lopezator/nexmo-go)](https://goreportcard.com/report/github.com/lopezator/nexmo-go)
[![GoDoc](https://godoc.org/github.com/lopezator/nexmo-go/go?status.svg)](https://godoc.org/github.com/lopezator/nexmo-go)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

# nexmo-go

Nexmo REST API client for Go.

Hugely inspired by: https://github.com/kevinburke/twilio-go

As it's a WIP it only supports SMS and TTS Call sending by now.

# Send a SMS

```
const apiKey = "my-api-key"
const apiSecret = "my-api-secret"

// Create a client
client := NewClient(apiKey, apiSecret, nil)

// Send a message
// Nexmo allows to use your either a random text as `from` value or your nexmo phone
msg, err := client.Messages.SendMessage("ME", "+34666666666", "Message sent via nexmo-go")
```


# Make a TTS Call

```
const apiKey = "my-api-key"
const apiSecret = "my-api-secret"

// Create a client
client := NewClient(apiKey, apiSecret, nil)

// Make a TTS call
msg, err := client.Calls.MakeTTSCall("+15111111111", "+34666666666", "TTS call sent via nexmo-go")
```