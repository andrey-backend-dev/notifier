package notify

import (
	"bytes"
	"log/slog"
	"testing"
)

func TestNameEmail(t *testing.T) {
	expected := "Email"

	if result := (EmailSender{}).Name(); expected != result {
		t.Errorf("Expected: %s, actual: %s", expected, result)
	}
}

func TestNameSms(t *testing.T) {
	expected := "Sms"

	if result := (SmsSender{}).Name(); expected != result {
		t.Errorf("Expected: %s, actual: %s", expected, result)
	}
}

func TestLoggingSenderSend(t *testing.T) {
	var buffer bytes.Buffer
	ls, _ := NewLoggingSender(EmailSender{}, *slog.New(slog.NewTextHandler(&buffer, nil)))
	n := Notification{}

	if err := ls.Send(n); err != nil {
		t.Fatal("Error should not be returned")
	}

	if buffer.Len() == 0 {
		t.Fatal("Buffer should be filled with log messages")
	}
}
