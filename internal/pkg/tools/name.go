package tools

import "strings"

func FormatTypeName(s string) string {
	name := strings.ToUpper(s[0:1])
	for i := 1; i < len(s); i++ {
		if s[i] == '_' {
			i++
			name += strings.ToUpper(s[i : i+1])
		} else {
			name += s[i : i+1]
		}
	}
	return name
}
