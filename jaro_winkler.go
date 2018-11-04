package main

var BoostThreshold float64 = 0.7

var PrefixLen = 4

func JaroWinkler(s1, s2 []rune) float64 {
	jaro := Jaro(s1, s2)
	if jaro < BoostThreshold {
		return jaro
	}
	var prefix int
	for i := 0; i < min(len(s1), len(s2)); i++ {
		if s1[i] == s2[i] {
			prefix++
			if prefix == PrefixLen {
				break
			}
			continue
		}
		break
	}
	return jaro + 0.1*float64(prefix)*(1-jaro)
}

func Jaro(s1, s2 []rune) float64 {
	if len(s1) == 0 && len(s2) == 0 {
		return 1
	}
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}
	matchDist := max(len(s1), len(s2))/2 - 1

	var short, long []rune
	if len(s1) > len(s2) {
		long, short = s1, s2
	} else {
		long, short = s2, s1
	}
	s1Matches := make([]bool, len(short))
	s2Matches := make([]bool, len(long))
	matches := 0.
	transposition := 0.
	for i := range short {
		start := max(0, i-matchDist)
		end := min(i+matchDist+1, len(long))
		for j := start; j < end; j++ {
			if s2Matches[j] {
				continue
			}
			if short[i] != long[j] {
				continue
			}
			s1Matches[i] = true
			s2Matches[j] = true
			matches++
			break
		}
	}
	if matches == 0 {
		return 0
	}
	var k int
	for i := range short {
		if !s1Matches[i] {
			continue
		}
		for !s2Matches[k] {
			k++
		}
		if short[i] != long[k] {
			transposition++
		}
		k++
	}

	return (matches/float64(len(s1)) +
		matches/float64(len(s2)) +
		(matches-transposition/2)/matches) / 3
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
