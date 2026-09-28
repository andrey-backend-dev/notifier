package notify

import "fmt"

type ValidationError struct {
	Field  string
	Reason string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("поле %s: %s", e.Field, e.Reason)
}

func checkRecipient(n Notification) error {
	if n.Recipient == "" {
		return ValidationError{Field: "recipient", Reason: "empty"}
	}
	return nil
}

func Validate(n Notification) error {
	return checkRecipient(n)
}
