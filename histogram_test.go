package histogram_test

import (
	"github.com/daneharrigan/histogram"
	"testing"
)

const SIZE = 1000

func TestMerge(t *testing.T) {
	h1 := histogram.New(10)
	h2 := histogram.New(10)

	for i := 0; i < 10; i++ {
		h1.Insert(float64(i))
	}

	for i := 0; i < 5; i++ {
		h2.Insert(float64(i))
	}

	assertEqual(t, 10, len(h1.Bins))
	assertEqual(t, 5, len(h2.Bins))

	h1.Merge(h2.Bins)
	assertEqual(t, 10, len(h1.Bins))

	for i := 0; i < 5; i++ {
		assertEqual(t, i, h1.Bins[i].Point)
		assertEqual(t, 2, h1.Bins[i].Count)
	}

	for i := 5; i < 10; i++ {
		assertEqual(t, i, h1.Bins[i].Point)
		assertEqual(t, 1, h1.Bins[i].Count)
	}
}

func BenchmarkInsert(b *testing.B) {
	h := histogram.New(SIZE)
	for i := 0; b.N > i; i++ {
		h.Insert(float64(i))
	}
}

func BenchmarkMerge(b *testing.B) {
	h1 := histogram.New(SIZE)
	h2 := histogram.New(SIZE)

	for i := 0; b.N > i; i++ {
		h1.Insert(float64(i))
		h2.Insert(float64(i))
	}

	h1.Merge(h2.Bins)
}

func assertEqual(t *testing.T, a, b interface{}) {
	var i, n float64

	switch a.(type) {
	case float64: i = a.(float64)
	case int: i = float64(a.(int))
	}

	switch b.(type) {
	case float64: n = b.(float64)
	case int: n = float64(b.(int))
	}

	if i != n {
		t.Fatalf("%f was not %f", i, n)
	}
}
