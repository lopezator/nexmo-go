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

	// Make a tts call
	_, err := client.Calls.MakeTTSCall("+15111111111", "+34666666666", "TTS call sent via nexmo-go", "5")
	if err != nil {
		log.Fatalf("couldn't make the tts call using nexmo-go: %v", err)
	}
	fmt.Println("tts call sent using nexmo-go!")
}
