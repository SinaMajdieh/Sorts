package Sorts

import "fmt"

type bubble struct {
	array     []float64
	ascending bool
	len       int
}

func (b *bubble) sort() {
	for i := 0; i < b.len; i++ {
		for j := 0; j < b.len-i; j++ {
			switch b.ascending {
			case true:
				if b.array[j] > b.array[j+1] {
					swap(j, j+1, b.array)
				}
			case false:
				if b.array[j] < b.array[j+1] {
					swap(j, j+1, b.array)
				}
			}
		}
	}
}

//Bubble Sort Algorithm
func BubbleSort(Ascending bool, Array interface{}) {
	b := bubble{}
	switch Array.(type) {
	case []int:
		intstofloats(Array.([]int), b.array)
	case []float64:
		b.array = Array.([]float64)
	default:
		fmt.Printf("Cannot sort type %T\n", Array)
		return
	}
	b.len = len(b.array) - 1
	b.ascending = Ascending

	b.sort()
}
