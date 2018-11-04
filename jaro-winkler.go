package main

import (
	"fmt"
)

var BoostThreshold float64 = 0.7

func main() {
	s1 := []rune("massey")
	s2 := []rune("massie")
	fmt.Printf("jaro winkler for %c and %c is %f", s1, s2, jaroWinkler(s1, s2))
}

func jaroWinkler(s1, s2 []rune) float64 {
	// l represents the max len
	l := max(len(s1), len(s2))
	t := make([]int, l)

	// The maximum allowed distance
	d := l/2 - 1

	var short, long []rune

	long = s2
	short = s1
	if len(s1) > len(s2) {
		long = s1
		short = s2
	}
	set2 := make(map[int]bool)
	for i := 0; i < len(short); i++ {
		for j := 0; j < len(long); j++ {
			// Only take those that are not visited yet.
			if set2[j] {
				continue
			}
			if short[i] == long[j] {
				set2[j] = true
				t[i] = j + 1
				// Avoid overlapping indices, since the alphabet can
				// appear later in the loop.
				break
			}
		}
	}
	var prefix, largest, tr, m int
	PREFIX := 4
	for _, v := range t {
		// Check for consecutive matches. Needs to start from 0.

		// The text is not found.
		if v == 0 {
			continue
		}
		if v-largest == 1 && prefix < PREFIX {
			prefix++
		}
		if abs(v-largest) <= d {
			m++
		}
		if v > largest {
			largest = v
		} else {
			tr++
		}
	}
	jaro := 1 / float64(3) * (float64(m)/float64(len(s1)) + float64(m)/float64(len(s2)) + float64(m-tr)/float64(m))
	if jaro < BoostThreshold {
		return jaro
	}
	// jaro winkler
	jw := jaro + 0.1*float64(prefix)*(1-jaro)
	return jw
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
