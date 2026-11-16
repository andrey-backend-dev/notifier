package notify

import "errors"

var (
	ErrNotFound error = errors.New("уведомление не найдено")

	ErrAlreadySent error = errors.New("уведомление уже отправлено")

	ErrRateLimited error = errors.New("превышен лимит запросов")
)
