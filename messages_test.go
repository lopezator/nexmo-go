package nexmo

import (
	"strconv"
	"testing"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()
	client, s := getServer(sendMessageResponse)
	defer s.Close()
	msg, err := client.Messages.SendMessage(from, to, "nexmo-go testing!")
	if err != nil {
		t.Fatal(err)
	}
	msgCount, err := strconv.Atoi(msg.MessageCount)
	if err != nil {
		t.Errorf("type conversion failed, messageCount is not an integer")
	}
	if msgCount <= 0 {
		t.Errorf("expected MessageCount to be greater than 0, got error")
	}
	status, err := strconv.Atoi(msg.Messages[0].Status)
	if err != nil {
		t.Errorf("type conversion failed, Status is not an integer")
	}
	if status != 0 {
		t.Errorf("expected status to be 0, got %s", msg.Messages[0].Status)
	}
}
