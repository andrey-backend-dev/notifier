package notify

import "testing"

func TestRetryHappyPath(t *testing.T) {
	notification := Notification{Status: StatusCreated}
	var retryTimes uint32 = 3

	for range retryTimes {
		notification.Retry()
	}

	if notification.Attempts != retryTimes {
		t.Errorf("Expected attempts: %d, actual: %d", retryTimes, notification.Attempts)
	}

	if notification.Status != StatusDelivering {
		t.Errorf("Expected status: %s, actual: %s", StatusDelivering, notification.Status)
	}
}

func TestRetryWhenStatusIsDelivered(t *testing.T) {
	status := StatusDelivered
	notification := Notification{Status: status}

	for range 3 {
		notification.Retry()
	}

	if notification.Attempts != 0 {
		t.Errorf("Expected attempts: %d, actual: %d", 0, notification.Attempts)
	}

	if notification.Status != status {
		t.Errorf("Expected status: %s, actual: %s", status, notification.Status)
	}
}

func TestString(t *testing.T) {
	notification := Notification{ID: 67, Recipient: "Vasyan", Body: "pizdets VASYA RUN VASYA RUN!!!", Status: StatusDelivering, Attempts: 3}
	expected := "notification#67 to Vasyan [delivering, attempt 3]"

	if result := notification.String(); expected != result {
		t.Errorf("Expected: %s, actual: %s", expected, result)
	}
}
