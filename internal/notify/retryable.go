package notify

import (
	"fmt"
	"time"
)

type RetryableError struct {
	Op    string
	After time.Duration
	Err   error
}

func NewRetryableError(after time.Duration, op string, err error) error {
	return &RetryableError{After: after, Op: op, Err: err}
}

func (e *RetryableError) Error() string {
	return fmt.Sprintf("%s: %v", e.Op, e.Err)
}

func (e *RetryableError) Unwrap() error {
	return e.Err
}
