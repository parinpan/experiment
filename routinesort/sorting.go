package routinesort

const (
	ASC = iota + 1
	DESC
)

type Number interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

func Sort[T Number](numbers []T, order int) []T {
	if len(numbers) == 0 {
		return []T{}
	}

	chamberOne := make(chan T)
	chamberTwo := make(chan T)

	done := make(chan struct{}, 2)
	sorted := make([]T, 0, len(numbers))

	go func() {
		for number := range chamberOne {
			if n := len(sorted); n == 0 || compare(number, sorted[n-1], order) {
				sorted = append(sorted, number)
			} else {
				curr := number
				last := sorted[n-1]

				go func() {
					chamberTwo <- curr
					chamberTwo <- last
				}()

				sorted = sorted[:n-1]
			}

			if len(sorted) == len(numbers) {
				done <- struct{}{}
				break
			}
		}
	}()

	go func() {
		for number := range chamberTwo {
			chamberOne <- number
		}
		done <- struct{}{}
	}()

	for i := 0; i < len(numbers); i++ {
		chamberTwo <- numbers[i]
	}

	<-done
	close(chamberTwo)
	<-done
	close(chamberOne)

	return sorted
}

func compare[T Number](a, b T, order int) bool {
	if order == ASC {
		return a >= b
	}
	return b >= a
}
