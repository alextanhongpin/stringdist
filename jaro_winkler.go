package stringdist

var BoostThreshold float64 = 0.7

var PrefixLen = 4.

func JaroWinkler(s1, s2 string) float64 {
	jaro := Jaro(s1, s2)
	if jaro < BoostThreshold {
		return jaro
	}
	prefix := 0.
	for i := 0; i < min(len(s1), len(s2)); i++ {
		if prefix == PrefixLen {
			break
		}
		if s1[i] == s2[i] {
			prefix++
		} else {
			break
		}
	}
	return jaro + 0.1*prefix*(1-jaro)
}

func Jaro(source, target string) float64 {
	// Should return 1 if both are empty.
	if len(source) == 0 && len(target) == 0 {
		return 1
	}
	// If either one is empty, the distance should be 0.
	if len(source) == 0 || len(target) == 0 {
		return 0
	}
	// Place the shorter string on the outer double loop.
	if len(source) < len(target) {
		return jaro(source, target, len(target))
	}
	return jaro(target, source, len(source))
}

func jaro(s, t string, longest int) float64 {
	matchDistance := longest/2 - 1
	matchSource := make([]bool, len(s))
	matchTarget := make([]bool, len(t))
	matches := 0.
	transpositions := 0.
	for i := range s {
		start := max(0, i-matchDistance)
		end := min(i+matchDistance+1, len(t))
		for j := start; j < end; j++ {
			if matchTarget[j] {
				continue
			}
			if s[i] != t[j] {
				continue
			}
			matchSource[i] = true
			matchTarget[j] = true
			matches++
			break
		}
	}
	if matches == 0 {
		return 0
	}
	var j int
	for i := range s {
		if !matchSource[i] {
			continue
		}
		for !matchTarget[j] {
			j++
		}
		if s[i] != t[j] {
			transpositions++
		}
		j++
	}
	return (matches/float64(len(s)) +
		matches/float64(len(t)) +
		(matches-transpositions/2)/matches) / 3
}
