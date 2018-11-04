package main

import (
	"math"
	"testing"
)

func TestJaroWinkler(t *testing.T) {
	tests := []struct {
		s1, s2        string
		jaro, winkler float64
	}{
		{"SHACKLEFORD", "SHACKELFORD", 0.970, 0.982},
		{"DUNNINGHAM", "CUNNIGHAM", 0.896, 0.896},
		{"NICHLESON", "NICHULSON", 0.926, 0.956},
		{"JONES", "JOHNSON", 0.790, 0.832},
		{"MASSEY", "MASSIE", 0.889, 0.933},
		{"ABROMS", "ABRAMS", 0.889, 0.922},
		{"HARDIN", "MARTINEZ", 0.722, 0.722},
		// {"HARDIN", "MARTINEZ", 0.000, 0.000},
		// {"ITMAN", "SMITH", 0.622, 0.622},
		{"ITMAN", "SMITH", 0.467, 0.467},

		{"JERALDINE", "GERALDINE", 0.926, 0.926},
		{"MARHTA", "MARTHA", 0.944, 0.961},
		{"MICHELLE", "MICHAEL", 0.869, 0.921},
		{"JULIES", "JULIUS", 0.889, 0.933},
		{"TANYA", "TONYA", 0.867, 0.880},
		{"DWAYNE", "DUANE", 0.822, 0.840},
		{"SEAN", "SUSAN", 0.783, 0.805},
		{"JON", "JOHN", 0.917, 0.933},
		// {"JON", "JAN", 0.000, 0.000},
		{"JON", "JAN", 0.778, 0.800},
	}
	for _, tt := range tests {
		jaro := Jaro([]rune(tt.s1), []rune(tt.s2))
		winkler := JaroWinkler([]rune(tt.s1), []rune(tt.s2))
		if math.Round(jaro*1000)/1000 != tt.jaro {
			t.Fatalf("expected jaro %f, got %f for %s and %s", tt.jaro, jaro, tt.s1, tt.s2)
		}
		if math.Round(winkler*1000)/1000 != tt.winkler {
			t.Fatalf("expected winkler %f, got %f for %s and %s", tt.winkler, winkler, tt.s1, tt.s2)
		}
	}

}
