package nexmo

import (
	"testing"
)

func TestSendCall(t *testing.T) {
	t.Parallel()
	client, s := getServer(sendCallResponse)
	defer s.Close()
	call, err := client.Calls.MakeTTSCall("", "", "nexmo-go testing!")
	if err != nil {
		t.Fatal(err)
	}
	if call.CallID == "" {
		t.Error("CallID should not be blank!")
	}
	if call.Status != "0" {
		t.Errorf("expected status to be 0, got %s", call.Status)
	}
	if call.ErrorText != "Success" {
		t.Errorf("expected direction to be Success, got %s", call.ErrorText)
	}
}
