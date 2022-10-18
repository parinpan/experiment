# RoutineSort Algorithm

RoutineSort is an experimental sorting algorithm, it utilizes Go channels to complete the sorting task. The idea is to have **two chambers (channels)** to hold numbers and place them in a correct order in a new array (not for production 😛). The algorithm:

1. Feed all unsorted array numbers to `chamberTwo`
2. A spawned go-routine consumes the `chamberTwo` to get numbers, numbers are then handed over to `chamberOne`
3. A spawned go-routine consumes the `chamberOne` to get numbers fed by `chamberTwo`
4. `chamberOne` go-routine will receive a number and check:
    - if new array is empty or last element in the new array is less than or equal to the number, then append the number to the array
    - else, push back the number to `chamberTwo` alongwith the last number in the array, then pop the last element in the array
5. Repeat step 2 to step 4 until array is sorted
6. Exit


## Usage

Install:
```
go get github.com/parinpan/experiment
```

Function:
```go
import (
   "github.com/parinpan/experiment/routinesort"
)

func main() {
    sortedASC := routinesort.Sort[int32]([]int32{5, 3, 0, -1, 5, 100, 3, 4, 7}, routinesort.ASC)
    sortedDESC := routinesort.Sort[float64]([]float64{5, 3, 0, -1, 5, 100, 3, 4, 7}, routinesort.DESC)
    
    fmt.Println(sortedASC) // output: []int32{-1, 0, 3, 3, 4, 5, 5, 7, 100}
    fmt.Println(sortedDESC) // output: []float64{100, 7, 5, 5, 4, 3, 3, 0, -1}
}
```
