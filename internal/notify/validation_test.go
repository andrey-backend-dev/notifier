package notify

import "testing"

func TestValidateSuccessful(t *testing.T) {
	n := Notification{Recipient: "Vasyan"}

	if err := Validate(n); err != nil {
		t.Fatalf("Error should not be returned: %v", err)
	}
}

func TestValidateErrorWhenRecipientIsEmpty(t *testing.T) {
	n := Notification{Recipient: ""}

	if err := Validate(n); err == nil {
		t.Fatal("Error should be returned")
	}
}
