# fuzzing - brown bag


# Steps TDD 

## Write unit test

```go
func TestReverse(t *testing.T) {
	t.Parallel()

	input := "nginx"
	want := "xnign"

	got := words.Reverse(input)
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
```

## Write a function to compile

```go
// Reverse takes a string and reverses it.
func Reverse(s string) string {
	return ""
}
```

## Run tests - watch it fail

## Write a function to pass the test

```go
// Reverse takes a string and reverses it.
func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
```

## Refactor tests

- add more tests to cover various inputs
- possible group test input data into ```table tests```

```go
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

```


## Run tests and watch it pass (green)

```
➜  fuzzing-demo git:(main) ✗ go test -v
=== RUN   TestReverse
=== PAUSE TestReverse
=== CONT  TestReverse
=== RUN   TestReverse/single_word
=== RUN   TestReverse/multiple_words
=== RUN   TestReverse/spaces_only
=== RUN   TestReverse/special_chars
--- PASS: TestReverse (0.00s)
    --- PASS: TestReverse/single_word (0.00s)
    --- PASS: TestReverse/multiple_words (0.00s)
    --- PASS: TestReverse/spaces_only (0.00s)
    --- PASS: TestReverse/special_chars (0.00s)
=== RUN   FuzzReverse
=== RUN   FuzzReverse/seed#0
=== RUN   FuzzReverse/seed#1
=== RUN   FuzzReverse/seed#2
=== RUN   FuzzReverse/seed#3
--- PASS: FuzzReverse (0.00s)
    --- PASS: FuzzReverse/seed#0 (0.00s)
    --- PASS: FuzzReverse/seed#1 (0.00s)
    --- PASS: FuzzReverse/seed#2 (0.00s)
    --- PASS: FuzzReverse/seed#3 (0.00s)
PASS
ok  	github.com/jjngx/words	0.262s
```

TIP:
```
go install github.com/bitfield/gotestdox/cmd/gotestdox@latest
```

```
➜  fuzzing-demo git:(main) ✗ gotestdox
github.com/jjngx/words:
 ✔ Reverse single word (0.00s)
 ✔ Reverse multiple words (0.00s)
 ✔ Reverse spaces only (0.00s)
 ✔ Reverse special chars (0.00s)
 ✔ Reverse (0.00s)
```


# Steps Fuzzing

```go
func FuzzReverse(f *testing.F) {
	// Initial slice od strings to "feed" to corpus
	inputs := []string{"ingress controller", " ", "!12345", "*.example.com"}

	// Build test corpus with initial input data
	for _, input := range inputs {
		f.Add(input)
	}

	// Fuzzing
	f.Fuzz(func(t *testing.T, input string) {
		firstReverse := words.Reverse(input)
		secondReverse := words.Reverse(firstReverse)

		// Assumption:
		// double reverse should produce the same string as the input

		if input != secondReverse {
			t.Errorf("want %q, got %q", input, secondReverse)
		}

		// validate if both: input and reversed strings are valid UTF-8
		if utf8.ValidString(input) && !utf8.ValidString(secondReverse) {
			t.Errorf("want valid utf8 string, got %q", secondReverse)
		}
	})
}

```



## Write fuzz test

- add a slice of inputs
- can't predict ```wants```
- decide under what conditions test should fail

## Run unit tests and fuzz tests

- run all tests without fuzzing engine

```
➜  fuzzing-demo git:(main) ✗ go test -v
=== RUN   TestReve
=== PAUSE TestReve
=== CONT  TestReve
=== RUN   TestReve/single_word
=== RUN   TestReve/multiple_words
=== RUN   TestReve/spaces_only
=== RUN   TestReve/special_chars
--- PASS: TestReve (0.00s)
    --- PASS: TestReve/single_word (0.00s)
    --- PASS: TestReve/multiple_words (0.00s)
    --- PASS: TestReve/spaces_only (0.00s)
    --- PASS: TestReve/special_chars (0.00s)
=== RUN   FuzzReverse
=== RUN   FuzzReverse/seed#0
=== RUN   FuzzReverse/seed#1
=== RUN   FuzzReverse/seed#2
=== RUN   FuzzReverse/seed#3
=== RUN   FuzzReverse/seed#4
--- PASS: FuzzReverse (0.00s)
    --- PASS: FuzzReverse/seed#0 (0.00s)
    --- PASS: FuzzReverse/seed#1 (0.00s)
    --- PASS: FuzzReverse/seed#2 (0.00s)
    --- PASS: FuzzReverse/seed#3 (0.00s)
    --- PASS: FuzzReverse/seed#4 (0.00s)
PASS
ok  	github.com/jjngx/words	0.096s
```


- run fuzz tests

```
➜  fuzzing-demo git:(main) ✗ go test --fuzz=FuzzReverse
fuzz: elapsed: 0s, gathering baseline coverage: 0/44 completed
fuzz: elapsed: 0s, gathering baseline coverage: 44/44 completed, now fuzzing with 10 workers
^Cfuzz: elapsed: 2s, execs: 817317 (480202/sec), new interesting: 0 (total: 44)

```

Fuzz tests with time limit
```bash
➜  fuzzing-demo git:(main) ✗ go test --fuzz=FuzzReverse --fuzztime 10s
fuzz: elapsed: 0s, gathering baseline coverage: 0/44 completed
fuzz: elapsed: 0s, gathering baseline coverage: 44/44 completed, now fuzzing with 10 workers
fuzz: elapsed: 3s, execs: 1414225 (471392/sec), new interesting: 0 (total: 44)
fuzz: elapsed: 6s, execs: 2979746 (521733/sec), new interesting: 0 (total: 44)
fuzz: elapsed: 9s, execs: 4606196 (542159/sec), new interesting: 0 (total: 44)
fuzz: elapsed: 10s, execs: 5097727 (449477/sec), new interesting: 0 (total: 44)
PASS
ok  	github.com/jjngx/words	10.206s
```

## Fix failing functions

