package notify

import (
	"errors"
	"log/slog"
	"time"
)

type SendFunc func(n Notification) error

func WithRetry(inner SendFunc, attempts int, latency time.Duration) SendFunc {
	if attempts < 1 {
		return func(n Notification) error {
			return errors.New("Количество попыток для retry-обертки должно быть > 0")
		}
	}

	return func(n Notification) (err error) {
		var i int

		defer func() {
			slog.Debug("Отправка завершена", "id", n.ID, "attempts", i+1, "err", err)
		}()

		for i = range attempts {
			if err = inner(n); err == nil {
				return nil
			}

			if attempts > i+1 {
				time.Sleep(latency)
			}
		}

		return err
	}
}
