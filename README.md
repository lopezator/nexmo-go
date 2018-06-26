[![Build Status](https://travis-ci.org/lopezator/nexmo-go.svg?branch=master)](https://travis-ci.org/lopezator/nexmo-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/lopezator/nexmo-go)](https://goreportcard.com/report/github.com/lopezator/nexmo-go)
[![GoDoc](https://godoc.org/github.com/lopezator/nexmo-go/go?status.svg)](https://godoc.org/github.com/lopezator/nexmo-go)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

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


