package words_test

import (
	"testing"

	"github.com/jjngx/words"
)

func TestReverse(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "single word",
			input: "nginx",
			want:  "xnign",
		},
		{
			name:  "multiple words",
			input: "hello from ingress-controller",
			want:  "rellortnoc-ssergni morf olleh",
		},
		{
			name:  "spaces only",
			input: "   ",
			want:  "   ",
		},
		{
			name:  "special chars",
			input: "!nginx123&",
			want:  "&321xnign!",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := words.Reverse(tc.input)
			if tc.want != got {
				t.Errorf("want %q, got%q", tc.want, got)
			}
		})
	}
}
