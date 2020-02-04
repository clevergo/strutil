package strutil

import "unicode"

// InitialToUpper
func InitialToUpper(s string) string {
	for i, v := range s {
		return string(unicode.ToUpper(v)) + s[i+1:]
	}

	return ""
}

// InitialToLower
func InitialToLower(s string) string {
	for i, v := range s {
		return string(unicode.ToLower(v)) + s[i+1:]
	}

	return ""
}
