package pad

import "testing"

var padTests = []struct {
	input string
	len   int
	want  string
}{
	{"foobar", 10, "****foobar"},
	{"foobarbarquux", 8, "foobarbarquux"},
	{"", 13, "*************"},
}

func TestLeft(t *testing.T) {
	for i, test := range padTests {
		output := Left(test.input, test.len)
		if output != test.want {
			t.Errorf("test %d failed: want %q, got %q", i, test.want, output)
		}
	}
}
