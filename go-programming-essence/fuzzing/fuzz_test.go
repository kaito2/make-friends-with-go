package fuzzing

import (
	"testing"
)

func FuzzTest(f *testing.F) {
	f.Add(8000, "Hello")
	f.Fuzz(func(t *testing.T, a int, b string) {
		target(a, b)
	})
}
