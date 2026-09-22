package notify

import (
	"bufio"
	"fmt"
	"os"
)

func WriteReport(path string, rows []string) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Создание отчёта %s: %w", path, err)
	}

	firstErrorFunc := func(inner func() error) {
		if err != nil {
			inner()
			return
		}

		err = inner()
	}

	defer firstErrorFunc(file.Close)

	writer := bufio.NewWriter(file)
	defer firstErrorFunc(writer.Flush)

	for _, row := range rows {
		if _, err = fmt.Fprint(writer, row); err != nil {
			break
		}
	}

	return err
}
