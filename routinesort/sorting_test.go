package routinesort

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
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

	t.Run("returns sorted array from @unpeacefulminds input suggestion", func(t *testing.T) {
		assert.EqualValues(t, []uint64{1, 3, 4, 5, 7}, Sort[uint64]([]uint64{1, 5, 3, 7, 4}, ASC))
		assert.EqualValues(t, []uint64{7, 5, 4, 3, 1}, Sort[uint64]([]uint64{1, 5, 3, 7, 4}, DESC))
	})

	t.Run("returns sorted array from @pc1ang input suggestion", func(t *testing.T) {
		assert.EqualValues(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, Sort[int]([]int{0, 9, 6, 2, 8, 4, 3, 7, 5, 1}, ASC))
		assert.EqualValues(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, Sort[int]([]int{0, 9, 6, 2, 8, 4, 3, 7, 5, 1}, DESC))
	})

	t.Run("returns sorted array with random numbers fed into the array", func(t *testing.T) {
		var numbers []int

		for i := 0; i < 1000; i++ {
			numbers = append(numbers, rand.Intn(10000))
		}

		sortedASC := Sort[int](numbers, ASC)
		sortedDESC := Sort[int](numbers, DESC)

		t.Run("on ASC", func(t *testing.T) {
			prev := sortedASC[0]

			for i := 1; i < len(sortedASC); i++ {
				assert.GreaterOrEqual(t, sortedASC[i], prev)
				prev = sortedASC[i]
			}
		})

		t.Run("on DESC", func(t *testing.T) {
			prev := sortedDESC[0]

			for i := 1; i < len(sortedDESC); i++ {
				assert.LessOrEqual(t, sortedDESC[i], prev)
				prev = sortedDESC[i]
			}
		})
	})
}
