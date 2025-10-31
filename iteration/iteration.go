package iteration

import "strings"

const repeatCount = 6

func Repeat(character string, count int) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
