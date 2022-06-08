package words

import (
	"errors"
	"unicode/utf8"
)

// Example: Reverse bytes
// Reverse takes a string and reverses it.
func ReverseString(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

// Example: Reverse runes
// Reverse takes a string and reverses it.
func ReverseRunes(s string) string {
	b := []rune(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

// Example: Reverse runes with UTF8 validation
// Reverse takes a string and reverses it.
func ReverseUTF8(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("invalid utf8 input")
	}
	b := []rune(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil
}
