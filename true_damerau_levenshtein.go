package stringdist

//func main() {
// gifts and profit is 5
// LD(CA,ABC) = 2
//fmt.Println("Hello, playground", trueDamerauLevenshtein("gifts", "profit"))
//}

type TrueDamerauLevenshtein struct {
	buffer [256]int
}

func NewTrueDamerauLevenshtein() *TrueDamerauLevenshtein {
	return &TrueDamerauLevenshtein{}
}

func (dl *TrueDamerauLevenshtein) Calculate(s, t string) int {
	m, n := len(s), len(t)

	// Initialize a new array the size of alphabet.
	da := dl.buffer
	for i := 0; i < len(da); i++ {
		da[i] = 0
	}

	// Initialize matrix d with the original length + 2.
	d := make([][]int, m+2)
	for i := 0; i < m+2; i++ {
		d[i] = make([]int, n+2)
	}
	maxdist := m + n
	d[0][0] = maxdist
	// m is inclusive.
	for i := 1; i <= m; i++ {
		d[i-1][0] = maxdist
		d[i-1][1] = i
	}
	for j := 1; j <= n; j++ {
		d[0][j-1] = maxdist
		d[1][j-1] = j
	}
	for i := 1; i <= m; i++ {
		db := 0
		for j := 1; j <= n; j++ {
			k := da[t[j-1]]
			l := db
			cost := 0
			if s[i-1] == t[j-1] {
				cost = 0
				db = j
			} else {
				cost = 1
			}
			k++
			l++
			d[i][j] = min(d[i-1][j-1]+cost, // Substitution.
				d[i][j-1]+1, // Insertion.
				d[i-1][j]+1, // Deletion.
				d[k-1][l-1]+(i-k-1)+1+(j-l-1)) // Transposition.

		}
		da[s[i-1]] = i
	}
	return d[m][n]
}
