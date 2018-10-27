package main

type EditDistanceCalculator interface {
	Compute(source, target []byte) int
}

type damerauLevenshtein struct {
	buffer [][]int
}

func NewDamerauLevenshtein(size int) *damerauLevenshtein {
	// Initialize reusable buffer. The size represents the maximum length
	// of the word allowed. 32 is a good size
	buffer := make([][]int, size)
	for i := 0; i < size; i++ {
		buffer[i] = make([]int, size)
	}
	return &damerauLevenshtein{
		buffer: buffer,
	}
}

func (d *damerauLevenshtein) Compute(s, t []byte) int {
	dp := d.buffer
	m, n := len(s), len(t)
	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			cost := 1
			if s[i-1] == t[j-1] {
				cost = 0
			}
			dp[i][j] = minimum(
				dp[i-1][j]+1,
				dp[i][j-1]+1,
				dp[i-1][j-1]+cost,
			)
			if i > 1 && j > 1 && s[i-1] == t[j-1] && s[i-2] == t[j-1] {
				dp[i][j] = minimum(
					dp[i][j],
					dp[i-2][j-2]+cost,
				)
			}
		}
	}
	return dp[m][n]
}

func minimum(head int, rest ...int) int {
	for _, i := range rest {
		if i < head {
			head = i
		}
	}
	return head
}

type levenshteinDistance struct {
	sourceBuffer []int
	targetBuffer []int
}

func NewLevenshteinDistance(size int) *levenshteinDistance {
	return &levenshteinDistance{
		sourceBuffer: make([]int, size),
		targetBuffer: make([]int, size),
	}
}

func (l *levenshteinDistance) Compute(s, t []byte) int {
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
