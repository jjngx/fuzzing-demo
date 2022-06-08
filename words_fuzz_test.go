package words_test

import (
	"testing"
	"unicode/utf8"

	"github.com/jjngx/words"
)

func FuzzReverseString(f *testing.F) {
	// Initial slice od strings to "feed" to corpus
	inputs := []string{"nginx", "! & $% ", "!12345", "*.example.com"}

	// Build test corpus with initial input data
	for _, input := range inputs {
		f.Add(input)
	}

	// Fuzzing
	f.Fuzz(func(t *testing.T, s string) {
		firstReverse := words.Reverse(s)
		secondReverse := words.Reverse(firstReverse)

		// Assumption:
		// double reverse should produce the same string as the input
		if s != secondReverse {
			t.Errorf("want %q, got %q", s, secondReverse)
		}
		// validate if both: input and reversed strings are valid UTF-8
		if utf8.ValidString(s) && !utf8.ValidString(secondReverse) {
			t.Errorf("want valid utf8 string, got %q", secondReverse)
		}
	})
}

func FuzzReverseUTF8(f *testing.F) {
	// Initial slice od strings to "feed" to corpus
	inputs := []string{"nginx", "! & $% ", "!12345", "*.example.com"}

	// Build test corpus with initial input data
	for _, input := range inputs {
		f.Add(input)
	}

	// Fuzzing
	f.Fuzz(func(t *testing.T, s string) {
		firstReverse, err := words.ReverseUTF8(s)
		if err != nil {
			return
		}
		secondReverse, err := words.ReverseUTF8(firstReverse)
		if err != nil {
			return
		}

		// Assumption:
		// double reverse should produce the same string as the input
		if s != secondReverse {
			t.Errorf("want %q, got %q", s, secondReverse)
		}
		// validate if both: input and reversed strings are valid UTF-8
		if utf8.ValidString(s) && !utf8.ValidString(secondReverse) {
			t.Errorf("want valid utf8 string, got %q", secondReverse)
		}
	})
}
