package util

import (
	"testing"
)

func assertArrayEquals(t *testing.T, got, expected []float64) {
	for i, v := range got {
		if expected[i] != v {
			t.Errorf("Got %f but was expecting %f", v, expected[i])
		}
	}
}

func TestCopy(t *testing.T) {

	t.Run("Test basic copy", func(t *testing.T) {
		original := []float64{1.0, 5.0}
		copy := Copy(original)
		assertArrayEquals(t, copy, original)
	})

	t.Run("Test deep copy", func(t *testing.T) {
		original := []float64{1.0, 5.0}
		copy := Copy(original)
		copy[0] = 0.0
		if copy[0] == original[0] {
			t.Errorf("Copy failed to produce a deep copy")
		}
	})

}

func TestMap(t *testing.T) {
	original := []float64{1.0, 5.0}
	expected := []float64{2.0, 10.0}
	got := Map(original, func(x float64) float64 { return 2 * x })
	assertArrayEquals(t, got, expected)
}

func TestMax(t *testing.T) {

	t.Run("Test basic max", func(t *testing.T) {
		original := []float64{1.0, 5.0}
		max := Max(original, 3)
		assertArrayEquals(t, max, []float64{3.0, 5.0})
	})

	t.Run("Test max doesn't change original", func(t *testing.T) {
		original := []float64{1.0, 5.0}
		Max(original, 3)
		assertArrayEquals(t, original, []float64{1.0, 5.0})
	})
}
func TestNVals(t *testing.T) {

	t.Run("Test normalizing array with NVals", func(t *testing.T) {
		original := []float64{1.0, 5.0}
		nVals := NVals(1/float64(len(original)), len(original))
		assertArrayEquals(t, nVals, []float64{0.5, 0.5})
	})
}

func TestAddTo(t *testing.T) {

	t.Run("Test add to", func(t *testing.T) {
		original := []float64{1.0, 5.0}
		AddTo(1, original)
		assertArrayEquals(t, original, []float64{2, 6})
	})
}

func TestMultBy(t *testing.T) {

	t.Run("Test mult by", func(t *testing.T) {
		original := []float64{1.0, 5.0}
		MultBy(2, original)
		assertArrayEquals(t, original, []float64{2, 10})
	})
}

func TestAddVectorTo(t *testing.T) {

	t.Run("Test AddVectorTo", func(t *testing.T) {
		original := []float64{1.0, 5.0}
		AddVectorTo([]float64{3.0, 4.0}, original)
		assertArrayEquals(t, original, []float64{4, 9})
	})

	t.Run("Test AddVectorTo twice", func(t *testing.T) {
		original := []float64{1.0, 5.0}
		AddVectorTo([]float64{3.0, 4.0}, original)
		AddVectorTo([]float64{10, 10}, original)
		assertArrayEquals(t, original, []float64{14, 19})
	})
}
