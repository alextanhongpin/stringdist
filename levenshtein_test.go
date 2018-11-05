package stringdist_test

import (
	"testing"

	"github.com/alextanhongpin/stringdist"
)

func TestLevenshtein(t *testing.T) {
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
	}
	levenshtein := stringdist.NewLevenshtein(32)
	for _, tt := range testcases {
		if res := levenshtein.Calculate(tt.source, tt.target); res != tt.score {
			t.Fatalf("expected %d, got %d for %s and %s", tt.score, res, tt.source, tt.target)
		}
	}
}
