package iteration

import "strings"

const RepeatedCount = 5

func Repeat(in string) string {
	var repeated strings.Builder
	for range RepeatedCount {
		repeated.WriteString(in)
	}
	return repeated.String()
}
