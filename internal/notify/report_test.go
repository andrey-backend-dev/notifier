package notify

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteReportHappyPath(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "testPath.txt")
	fileRows := []string{"hi!\n", "my name is andrey\n", "and, as you can see, i'm teaching Go"}
	expectedJoin := strings.Join(fileRows, "")

	if err := WriteReport(filePath, fileRows); err != nil {
		t.Fatalf("Error occurred: %v", err)
	}

	if content, err := extractContentFromFile(filePath); err != nil {
		t.Fatalf("Error occurred: %v", err)
	} else if expectedJoin != content {
		t.Fatalf("Expected: %s, actual: %s", expectedJoin, content)
	}
}

func extractContentFromFile(path string) (content string, err error) {
	var byteContent []byte

	if byteContent, err = os.ReadFile(path); err != nil {
		return "", err
	}

	return string(byteContent), err
}

func TestReport(t *testing.T) {
	notification1 := Notification{Channel: "email", Status: StatusUnidentified}
	notification2 := Notification{Channel: "push", Status: StatusDelivered}
	notification3 := Notification{Channel: "sms", Status: StatusCreated}
	notification4 := Notification{Channel: "sms", Status: StatusCreated}
	notification5 := Notification{Channel: "push", Status: StatusFailed}
	list := []Notification{notification1, notification2, notification3, notification4, notification5}
	expected := `email     unidentified=1    
push      delivered=1    failed=1    
sms       created=2    
`

	if result := Report(list); result != expected {
		t.Fatalf("Expected: \n%s, actual: \n%s", expected, result)
	}
}
