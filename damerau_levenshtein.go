package stringdist

type DamerauLevenshtein struct {
	buffer [][]int
}

func NewDamerauLevenshtein(size int) *DamerauLevenshtein {
	// Initialize reusable buffer. The size represents the maximum length
	// of the word allowed. 32 is a good size.
	buffer := make([][]int, size)
	for i := 0; i < size; i++ {
		buffer[i] = make([]int, size)
	}
	return &DamerauLevenshtein{
		buffer: buffer,
	}
}

func (d *DamerauLevenshtein) Calculate(s, t string) int {
	dp := d.buffer
	m, n := len(s), len(t)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	if max(m, n) > len(dp) {
		panic("length exceeded")
	}
	// Set the first column for each row equal to the row number.
	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	// Set the first row to equal the column number.
	for i := 0; i < n+1; i++ {
		dp[0][i] = i
	}
	// Starts from i = 1, which is an empty string.
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// Cost is 1 if the strings are not equal (requires
			// transposition), else 0.
			cost := 1
			if s[i-1] == t[j-1] {
				cost = 0
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
