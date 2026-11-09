package notify

import (
	"slices"
	"testing"
)

func TestDedupRecipients(t *testing.T) {
	initSlice := []string{"user", "user2", "user", "user4", "user4", "user3"}
	initSlicesCopy := slices.Clone(initSlice)
	expected := []string{"user", "user2", "user4", "user3"}

	if result := DedupRecipients(initSlice); !slices.Equal(expected, result) {
		t.Fatalf("Expected: %v, actual: %v", expected, result)
	}

	if !slices.Equal(initSlice, initSlicesCopy) {
		t.Fatalf("Expected init slice not to be modified")
	}
}
