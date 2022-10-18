package routinesort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	t.Run("returns empty output on an empty input", func(t *testing.T) {
		assert.Empty(t, Sort[float64]([]float64{}, ASC))
		assert.Empty(t, Sort[float64]([]float64{}, DESC))
	})

	t.Run("returns sorted if already sorted", func(t *testing.T) {
		assert.EqualValues(t, []int{1, 2}, Sort[int]([]int{1, 2}, ASC))
		assert.EqualValues(t, []int{100}, Sort[int]([]int{100}, ASC))
		assert.EqualValues(t, []float32{2, 1}, Sort[float32]([]float32{2, 1}, DESC))
		assert.EqualValues(t, []float32{100}, Sort[float32]([]float32{100}, DESC))
	})

	t.Run("returns sorted even if array items are not sorted", func(t *testing.T) {
		assert.EqualValues(t,
			[]int{-100, -1, 0, 0, 1, 1, 2, 3, 4, 4, 4, 99},
			Sort[int]([]int{1, 3, 2, 4, 0, -1, 99, -100, 4, 4, 1, 0}, ASC))

		assert.EqualValues(t, []float32{99, 4, 4, 4, 3, 2, 1, 1, 0, 0, -1, -100},
			Sort[float32]([]float32{1, 3, 2, 4, 0, -1, 99, -100, 4, 4, 1, 0}, DESC))

		assert.EqualValues(t, []uint64{0, 1, 2, 3, 4}, Sort[uint64]([]uint64{4, 3, 2, 1, 0}, ASC))
		assert.EqualValues(t, []uint{4, 3, 2, 1, 0}, Sort[uint]([]uint{0, 1, 2, 3, 4}, DESC))
	})
}
