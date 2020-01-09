package Sorts

import "fmt"

type insertion struct {
	array     []float64
	ascending bool
}

//Insertion Sort Algorithm
func (in *insertion) sort() {
	var j int
	var key float64
	for i := 0; i < len(in.array); i++ {
		key = in.array[i]
		j = i - 1
		switch in.ascending {
		case true:
			for j >= 0 && in.array[j] > key {
				in.array[j+1] = in.array[j]
				j--
			}
		case false:
			for j >= 0 && in.array[j] < key {
				in.array[j+1] = in.array[j]
				j--
			}
		}
		in.array[j+1] = key
	}
}

func InsertionSort(Ascending bool, Array interface{}) []float64 {
	in := insertion{}
	switch Array.(type) {
	case []int:
		array := intstofloats(Array.([]int))
		in.array = make([]float64, len(Array.([]int)))
		copy(in.array, array)
	case []float64:
		in.array = make([]float64, len(Array.([]float64)))
		copy(in.array, Array.([]float64))
	default:
		fmt.Printf("Cannot sort type %T\n", Array)
		return nil
	}

	in.ascending = Ascending
	in.sort()
	return in.array
}
