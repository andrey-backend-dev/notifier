package notify

import "testing"

func TestTruncate(t *testing.T) {
	input := "привет"
	expected := "прив…"
	maxSymbols := 5

	result, err := Truncate(input, maxSymbols)

	if err != nil {
		t.Errorf("Error should not be returned: %q", err)
	}

	if expected != result {
		t.Errorf("Expected: %s, actual: %s", expected, result)
	}
}

func TestTruncateWhenMaxSymbolsEqualsSymbolCountInInput(t *testing.T) {
	input := "привет"
	maxSymbols := 6

	result, err := Truncate(input, maxSymbols)

	if err != nil {
		t.Errorf("Error should not be returned: %q", err)
	}

	if input != result {
		t.Errorf("Expected: %s, actual: %s", input, result)
	}
}

func TestTruncateWhenMaxSymbolsMoreThanInInput(t *testing.T) {
	input := "привет"
	maxSymbols := 7

	result, err := Truncate(input, maxSymbols)

	if err != nil {
		t.Errorf("Error should not be returned: %q", err)
	}

	if input != result {
		t.Errorf("Expected: %s, actual: %s", input, result)
	}
}

func TestTruncateWhenMaxSymbolsIs1(t *testing.T) {
	input := "привет"
	expected := "…"
	maxSymbols := 1

	result, err := Truncate(input, maxSymbols)

	if err != nil {
		t.Errorf("Error should not be returned: %q", err)
	}

	if expected != result {
		t.Errorf("Expected: %s, actual: %s", expected, result)
	}
}

func TestTruncateWhenMaxSymbolsIs0(t *testing.T) {
	input := "привет"
	expected := ""
	maxSymbols := 0

	result, err := Truncate(input, maxSymbols)

	if err == nil {
		t.Errorf("Error should be returned")
	}

	if expected != result {
		t.Errorf("Expected: %s, actual: %s", expected, result)
	}
}

func TestTruncateWhenMaxSymbolsIsBelowZero(t *testing.T) {
	input := "привет"
	expected := ""
	maxSymbols := -1

	result, err := Truncate(input, maxSymbols)

	if err == nil {
		t.Errorf("Error should be returned")
	}

	if expected != result {
		t.Errorf("Expected: %s, actual: %s", expected, result)
	}
}
