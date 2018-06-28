package main

import (
	"fmt"
	"log"

	"github.com/lopezator/nexmo-go"
)

func main() {
	const apiKey = "my-api-key"
	const apiSecret = "my-api-secret"

	// Create a client
	client := nexmo.NewClient(apiKey, apiSecret, nil)

	// Send a message
	// Nexmo allows to use your either a random text as `from` value or your nexmo phone
	_, err := client.Messages.SendMessage("ME", "+34666666666", "Message sent via nexmo-go")
	if err != nil {
		log.Fatalf("couldn't send the message using nexmo-go: %v", err)
	}
	fmt.Println("message sent using nexmo-go!")
}
