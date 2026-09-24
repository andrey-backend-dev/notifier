package notify

import "testing"

func TestSafeHandleHappyPath(t *testing.T) {
	notification := Notification{}
	sendFunc := func(n Notification) error { return nil }

	if err := SafeHandle(notification, sendFunc); err != nil {
		t.Fatalf("Error occurred: %v", err)
	}
}

func TestSafeHandleWhenPanicOccurres(t *testing.T) {
	notification := Notification{}
	sendFunc := func(n Notification) error { panic("PANIC!") }
	err := SafeHandle(notification, sendFunc)

	if err == nil {
		t.Fatalf("Error should be returned, but not")
	}
}
