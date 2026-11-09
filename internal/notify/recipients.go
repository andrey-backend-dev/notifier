package notify

import "slices"

func DedupRecipients(list []string) []string {
	uniques := make(map[string]struct{}, len(list))
	out := make([]string, 0, len(list))

	for _, str := range list {
		if _, ok := uniques[str]; !ok {
			out = append(out, str)
			uniques[str] = struct{}{}
		}
	}

	return slices.Clip(out)
}
