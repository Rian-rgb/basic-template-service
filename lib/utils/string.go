package utils

import (
	"strings"
	"unicode"
)

func ToSnakeCase(value string) string {
	var result strings.Builder

	for i, r := range value {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteByte('_')
			}

			result.WriteRune(unicode.ToLower(r))
			continue
		}

		result.WriteRune(r)
	}

	return result.String()
}
