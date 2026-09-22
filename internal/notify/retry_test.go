package notify

import (
	"errors"
	"log/slog"
	"testing"
	"time"
)

func TestWithRetry(t *testing.T) {
	go slog.SetLogLoggerLevel(slog.LevelDebug)

	calls := 0
	flaky := func(n Notification) error {
		calls++
		if calls < 3 {
			return errors.New("провайдер недоступен")
		}
		return nil
	}

	send := WithRetry(flaky, 3, time.Millisecond)
	if err := send(Notification{ID: 1}); err != nil {
		t.Fatalf("ожидали успех с третьей попытки, получили %v", err)
	}
	if calls != 3 {
		t.Errorf("попыток %d, ожидали 3", calls)
	}
}

func TestWithRetryWhenAttemptsBelowOne(t *testing.T) {
	go slog.SetLogLoggerLevel(slog.LevelDebug)

	send := WithRetry(func(n Notification) error { return nil }, 0, time.Millisecond)

	if err := send(Notification{ID: 1}); err == nil {
		t.Fatalf("должна была вернуться ошибка")
	}
}
