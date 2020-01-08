package Sorts

import "fmt"

type quick struct {
	array     []float64
	ascending bool
}

func (q quick) partition(low int, high int) (partition int) {
	partition = low
	great := q.array[high]
	for i := low; i < high; i++ {
		switch q.ascending {
		case true:
			if q.array[i] <= great {
				swap(i, partition, q.array)
				partition++
			}
		case false:
			if q.array[i] >= great {
				swap(i, partition, q.array)
				partition++
			}
		}
	}
	swap(partition, high, q.array)
	return
}

func (q quick) sort(low int, high int) {
	if low < high {
		p := q.partition(low, high)
		q.sort(low, p-1)
		q.sort(p+1, high)
	}
}

//Quick Sort Algorithm
func QuickSort(Ascending bool, Array interface{}) {
	q := quick{}
	switch Array.(type) {
	case []int:
		intstofloats(Array.([]int) , q.array)
	case []float64:
		q.array = Array.([]float64)
	default:
		fmt.Printf("Cannot sort type %T\n", Array)
	}

	q.ascending = Ascending
	q.sort(0, len(q.array)-1)
}
