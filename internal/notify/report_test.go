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
