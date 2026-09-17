package notify

import "testing"

func TestParseStatus(t *testing.T) {
	expected := StatusDelivered

	if result, err := ParseStatus("delivered"); err != nil || expected != result {
		if err != nil {
			t.Errorf("Error should not be returned: %q", err)
		}
		t.Errorf("Expected: %v, actual: %v", expected, result)
	}
}

func TestParseStatusUnidentified(t *testing.T) {
	expected := StatusUnidentified

	if result, err := ParseStatus("some wrong status"); err == nil || expected != result {
		if err == nil {
			t.Error("Error should be returned")
		}
		t.Errorf("Expected: %v, actual: %v", expected, result)
	}
}

func TestStatusIsValid(t *testing.T) {
	expected := true
	mockStatus := StatusCreated

	if result := mockStatus.Valid(); expected != result {
		t.Errorf("Expected: %v, actual: %v", expected, result)
	}
}

func TestStatusIsNotValid(t *testing.T) {
	expected := false
	mockStatus := Status(100)

	if result := mockStatus.Valid(); expected != result {
		t.Errorf("Expected: %v, actual: %v", expected, result)
	}
}

func TestStatusIsNotValidWhenUnidentified(t *testing.T) {
	expected := false
	mockStatus := StatusUnidentified

	if result := mockStatus.Valid(); expected != result {
		t.Errorf("Expected: %v, actual: %v", expected, result)
	}
}
