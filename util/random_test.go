package util

import "testing"

func TestRandInt(t *testing.T) {

	min, max := 1, 10
	for i := 0; i < 1000; i++ {
		n := RandInt(min, max)
		if n < min || n > max {
			t.Fatalf("RandInt(%d, %d) = %d; want between %d and %d", min, max, n, min, max)
		}
	}
}
