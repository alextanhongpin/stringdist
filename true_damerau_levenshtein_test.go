package stringdist_test

import (
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"testing/quick"

	"github.com/alextanhongpin/stringdist"
)

type TestString struct {
	source string
	target string
}

func (TestString) Generate(r *rand.Rand, size int) reflect.Value {
	p := TestString{}
	a, b := rand.Int()%256, rand.Int()%256
	var sb strings.Builder
	var alphabets = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for i := 0; i < a; i++ {
		sb.WriteString(alphabets[i%26])
	}
	p.source = sb.String()
	sb.Reset()
	for i := 0; i < b; i++ {
		sb.WriteString(alphabets[i%26])
	}
	p.target = sb.String()

	return reflect.ValueOf(p)
}

// Test if the values are within the upper and lower boundary.
func TestTrueDamerauLevenshteinQuickCheck(t *testing.T) {
	t.Skip()
	T := t
	dl := stringdist.NewTrueDamerauLevenshtein()
	f := func(ts TestString) bool {
		s, t := ts.source, ts.target
		m, n := len(s), len(t)
		// At least the difference of the size of the two string
		lower := max(m, n) - min(m, n)
		// At most the max of the longer string.
		upper := max(m, n)

		dist := dl.Calculate(s, t)
		T.Log(dist, s, t, lower, upper)
		return dist >= lower && dist <= upper
	}
	if err := quick.Check(f, nil); err != nil {
		t.Fatal(err)
	}
}

func TestTrueDamerauLevenshteinDistance(t *testing.T) {
	dl := stringdist.NewTrueDamerauLevenshtein()
	tests := []struct {
		source, target string
		dist           int
	}{
		{"a", "", 1},
		{"", "a", 1},
		{"ac", "ca", 2},
		{"CA", "ABC", 3},
		{"4XHYWD", "YLKTW9", 5},
		{"YLKTW9", "4XHYWD", 5},
		{"hello", "hello", 0},
		{"kitten", "sitting", 3},
		{"", "xyz", 3},
		{"x", "xyz", 2},
		{"xy", "xyz", 1},
		{"yx", "xyz", 2},
		{"car", "rac", 2},
	}
	for _, tt := range tests {

		if dist := dl.Calculate(tt.source, tt.target); dist != tt.dist {
			t.Fatalf("expected %d, got %d for source %s, target %s", tt.dist, dist, tt.source, tt.target)
		}
	}
}
