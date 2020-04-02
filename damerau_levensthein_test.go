package stringdist_test

import (
	"log"
	"testing"
	"testing/quick"

	"github.com/alextanhongpin/stringdist"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Test if the values are within the upper and lower boundary.
func TestDamerauLevenshteinQuickCheck(t *testing.T) {
	threshold := 32
	dl := stringdist.NewDamerauLevenshtein(threshold)
	f := func(s, t string) bool {
		m, n := len(s), len(t)
		// At least the difference of the size of the two string
		lower := max(m, n) - min(m, n)
		// At most the max of the longer string.
		upper := max(m, n)

		// Can't measure more than the given threshold.
		if upper >= threshold {
			return true
		}

		dist := dl.Calculate(s, t)
		return dist >= lower && dist <= upper
	}
	if err := quick.Check(f, nil); err != nil {
		log.Fatal(err)
	}
}

func TestDamerauLevenshtein(t *testing.T) {
	testcases := []struct {
		source, target string
		score          int
	}{
		// The similarity between a string and itself is 0.
		{"", "", 0},
		{"a", "a", 0},
		// One edit operation.
		{"a", "", 1},
		// x, y and y, x should produce the same result.
		{"kitten", "sitting", 3},
		{"sitting", "kitten", 3},
		{"hello", "hello", 0},
		{"", "", 0},
		{"car", "rac", 2},
		{"4XHYWD", "YLKTW9", 5},
		{"YLKTW9", "4XHYWD", 5},
		{"CA", "ABC", 3},
	}
	levenshtein := stringdist.NewDamerauLevenshtein(32)
	for _, tt := range testcases {
		if res := levenshtein.Calculate(tt.source, tt.target); res != tt.score {
			t.Fatalf("expected %d, got %d for %s and %s", tt.score, res, tt.source, tt.target)
		}
	}
}
