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
func TestDamerauLevenshtein(t *testing.T) {
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
