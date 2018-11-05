package stringdist

/*
An optimized version of the levenshtein distance calculation.
*/

// Levenshtein represents the levenshtein operation.
type Levenshtein struct {
	// maxLen represents the maximum length of the strings to compute.
	// While it is not necessary to supply a value, this will help with the
	// performance when dealing with a lot of strings.
	maxLen       int
	sourceBuffer []int
	targetBuffer []int
}

// NewLevenshtein returns a new levenshtein operation.
func NewLevenshtein(size int) *Levenshtein {
	return &Levenshtein{
		maxLen:       size,
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

	m, n := len(s), len(t)
	v0 := l.sourceBuffer[:n+1]
	v1 := l.targetBuffer[:n+1]

	for i := 0; i < n; i++ {
		v0[i] = i
	}

	for i := 0; i < m; i++ {
		v1[0] = i + 1
		for j := 0; j < n; j++ {
			deletionCost := v0[j+1] + 1
			insertionCost := v1[j] + 1
			substitutionCost := v0[j] + 1
			if s[i] == t[j] {
				substitutionCost = v0[j]
			}
			v1[j+1] = min(deletionCost, insertionCost, substitutionCost)
		}
		v0, v1 = v1, v0
	}
	return v0[n]
}
