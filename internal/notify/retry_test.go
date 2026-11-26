package notify

import (
	"errors"
	"log/slog"
	"testing"
	"time"
)

var (
	firstError  = errors.New("первая ошибка")
	secondError = errors.New("вторая ошибка")
	thirdError  = errors.New("третья ошибка")
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
	if err := send(Notification{ID: 1}); err == nil {
		t.Fatalf("ожидали падение с первой попытки")
	}
	if calls != 1 {
		t.Errorf("попыток %d, ожидали 1", calls)
	}
}

func TestWithRetryRetryableError(t *testing.T) {
	calls := 0
	flaky := func(n Notification) error {
		calls++
		if calls == 1 {
			return NewRetryableError(1, "тестовая операция №1", firstError)
		}
		if calls == 2 {
			return NewRetryableError(2, "тестовая операция №2", secondError)
		}
		if calls == 3 {
			return NewRetryableError(1, "тестовая операция №3", thirdError)
		}
		return nil
	}

	send := WithRetry(flaky, 3, time.Millisecond)
	errs := send(Notification{ID: 1})

	if errs == nil {
		t.Fatalf("ожидали падение с повторяемыми ошибками")
	}

	if !errors.Is(errs, firstError) {
		t.Errorf("ожидали возникновение ошибки №1")
	}
	if !errors.Is(errs, secondError) {
		t.Errorf("ожидали возникновение ошибки №2")
	}
	if !errors.Is(errs, thirdError) {
		t.Errorf("ожидали возникновение ошибки №3")
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
