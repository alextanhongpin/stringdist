package stringdist

/*
An optimized version of the levenshtein distance calculation.
https://en.wikipedia.org/wiki/Levenshtein_distance#:~:text=contains%20the%20answer.-,Iterative%20with%20two%20matrix%20rows,-%5Bedit%5D
*/

// Levenshtein represents the levenshtein operation.
type Levenshtein struct {
	sourceBuffer []int
	targetBuffer []int
}

// NewLevenshtein returns a new levenshtein operation.
func NewLevenshtein(size int) *Levenshtein {
	return &Levenshtein{
		sourceBuffer: make([]int, size),
		targetBuffer: make([]int, size),
	}
}

// Calculate the levenshtein distance for the given source and target string.
func (l *Levenshtein) Calculate(s, t string) int {
	if len(s) == 0 {
		return len(t)
	}
	if len(t) == 0 {
		return len(s)
	}

	// It turns out that only two rows of the table – the previous row and the
	// current row being calculated – are needed for the construction, if one
	// does not want to reconstruct the edited input strings.
	m, n := len(s), len(t)

	if c, ok := resize(cap(l.sourceBuffer), n+1); ok {
		l.sourceBuffer = make([]int, c)
		l.targetBuffer = make([]int, c)
	}

	// Create two vectors of integer distances.
	v0 := l.sourceBuffer[:n+1]
	v1 := l.targetBuffer[:n+1]

	// Initialize v0 (the previous row of distances)
	// This row is A[0][i]: edit distance from an empty s to t;
	// That distance is the number of characters to append to s to make t.
	for i := range n {
		v0[i] = i
	}

	for i := range m {
		// Calculate v1 (current row distances) from the previous row v0.

		// First element of v1 is A[i + 1][0]
		//   edit distance is delete (i + 1) chars from s to match empty t
		v1[0] = i + 1

		// Use formula to fill in the rest of the row.
		for j := range n {
			// Calculating costs for A[i + 1][j + 1]
			deletionCost := v0[j+1] + 1
			insertionCost := v1[j] + 1
			substitutionCost := 0
			if s[i] == t[j] {
				substitutionCost = v0[j]
			} else {
				substitutionCost = v0[j] + 1
			}
			v1[j+1] = min(deletionCost, insertionCost, substitutionCost)
		}
		// Copy v1 (current row) to v0 (previous row) for next iteration
		// since data in v1 is always invalidated, a swap without copy would be more efficient.
		v0, v1 = v1, v0
	}

	// After the last swap, the results of v1 are now in v0.
	return v0[n]
}

func resize(c, n int) (int, bool) {
	if c < n {
		// Double the capacity.
		return n / c * c * 2, true
	}

	return 0, false
}
