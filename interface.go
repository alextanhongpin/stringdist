package stringdist

// Calculator calculates the distance between two string.
type Calculator interface {
	Calculate(source, target string) int
}
