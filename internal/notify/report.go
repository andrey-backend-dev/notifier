package notify

import (
	"bufio"
	"cmp"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
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

func Report(list []Notification) string {
	channelToStatusCountMap := make(map[string]map[Status]int)

	for _, notification := range list {
		if channelToStatusCountMap[notification.Channel] == nil {
			channelToStatusCountMap[notification.Channel] = make(map[Status]int)
		}
		channelToStatusCountMap[notification.Channel][notification.Status]++
	}

	var sb strings.Builder
	sortedChannelKeys := slices.Sorted(maps.Keys(channelToStatusCountMap))

	for _, channelKey := range sortedChannelKeys {
		fmt.Fprintf(&sb, "%-9s ", channelKey)

		statusKeys := slices.SortedFunc(maps.Keys(channelToStatusCountMap[channelKey]), func(a, b Status) int { return cmp.Compare(a.String(), b.String()) })

		for _, statusKey := range statusKeys {
			fmt.Fprintf(&sb, "%s=%-4d ", statusKey, channelToStatusCountMap[channelKey][statusKey])
		}

		sb.WriteByte('\n')
	}

	return sb.String()
}
