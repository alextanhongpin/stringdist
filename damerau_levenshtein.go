package stringdist

// The Damerau-Levenshtein distance measures the minimum number of
// single-character edits—insertions, deletions, substitutions, or
// transpositions of adjacent characters—needed to transform one string into
// another. Unlike standard Levenshtein distance, it correctly handles
// transpositions as a single operation (e.g., "CA" to "ABC" is 2, not 3)
type DamerauLevenshtein struct {
	buffer [][]int
}

func NewDamerauLevenshtein(size int) *DamerauLevenshtein {
	// Initialize reusable buffer. The size represents the maximum length
	// of the word allowed. 32 is a good size.
	buffer := make([][]int, size)
	for i := range size {
		buffer[i] = make([]int, size)
	}
	return &DamerauLevenshtein{
		buffer: buffer,
	}
}

func (d *DamerauLevenshtein) Calculate(s, t string) int {
	m, n := len(s), len(t)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	cm, mok := resize(cap(d.buffer), m+1)
	cn, nok := resize(cap(d.buffer[0]), n+1)
	switch {
	case mok:
		// If we recreate buffer m, we need to create buffer n too.
		d.buffer = make([][]int, cm)
		for i := range d.buffer {
			d.buffer[i] = make([]int, n+1)
		}
	case nok:
		// Only resize n.
		for i := range d.buffer {
			d.buffer[i] = make([]int, cn)
		}
	}

	dp := d.buffer
	dp = dp[:m+1]
	for i := range dp {
		dp[i] = dp[i][:n+1]
	}

	// Set the first column for each row equal to the row number.
	for i := range m + 1 {
		dp[i][0] = i
	}
	// Set the first row to equal the column number.
	for i := range n + 1 {
		dp[0][i] = i
	}
	// Starts from i = 1, which is an empty string.
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// Cost is 1 if the strings are not equal (requires
			// transposition), else 0.
			var cost int
			if s[i-1] == t[j-1] {
				cost = 0
			} else {
				cost = 1
			}
			// Find the minimum of the operations.
			dp[i][j] = min(
				dp[i-1][j]+1,      // Deletion
				dp[i][j-1]+1,      // Insertion
				dp[i-1][j-1]+cost, // Transposition
			)
			if i > 1 && j > 1 && s[i-1] == t[j-2] && s[i-2] == t[j-1] {
				dp[i][j] = min(
					dp[i][j],
					dp[i-2][j-2]+1,
				)
			}
		}
	}
	return dp[m][n]
}
