package iteration

import "strings"

func Repeat(in string, repeatedCount int) string {
	var repeated strings.Builder
	for range repeatedCount {
		repeated.WriteString(in)
	}
	return repeated.String()
}
