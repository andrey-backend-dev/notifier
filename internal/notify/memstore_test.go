package notify

import (
	"errors"
	"strings"
	"testing"
)

func TestMemStoreLoadMissing(t *testing.T) {
	s := NewMemStore()

	_, err := s.Load(42)

	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("ожидали ErrNotFound, получили %v", err)
	}
	if !strings.Contains(err.Error(), "42") {
		t.Errorf("в сообщении нет идентификатора: %v", err)
	}
}
