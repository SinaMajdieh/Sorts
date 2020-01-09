package Sorts

import "fmt"

type merge struct {
	array     []float64
	ascending bool
}

func (m *merge) sort(l int, mid int, r int) {
	n1 := mid - l + 1
	n2 := r - mid

	L := make([]float64, n1)
	R := make([]float64, n2)

	for i := 0; i < n1; i++ {
		L[i] = m.array[l+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = m.array[mid+1+i]
	}
	i, j, k := 0, 0, l
	for i < n1 && j < n2 {
		switch m.ascending {
		case true:
			if L[i] <= R[j] {
				m.array[k] = L[i]
				i++
			} else {
				m.array[k] = R[j]
				j++
			}
		case false:
			if L[i] >= R[j] {
				m.array[k] = L[i]
				i++
			} else {
				m.array[k] = R[j]
				j++
			}
		}
		k++
	}

	for i < n1 {
		m.array[k] = L[i]
		i++
		k++
	}

	for j < n2 {
		m.array[k] = R[j]
		j++
		k++
	}
}

func (m merge) mergesort(l int, r int) {
	if l < r {
		mid := l + (r-l)/2
		m.mergesort(l, mid)
		m.mergesort(mid+1, r)

		m.sort(l, mid, r)
	}
}

//Merge Sort Algorithm
func MergeSort(Ascending bool, Array interface{}) []float64 {
	m := merge{}
	switch Array.(type) {
	case []int:
		array := intstofloats(Array.([]int))
		m.array = make([]float64, len(Array.([]int)))
		copy(m.array, array)
	case []float64:
		m.array = make([]float64, len(Array.([]float64)))
		copy(m.array, Array.([]float64))
	default:
		fmt.Printf("Cannot sort type %T\n", Array)
		return nil
	}
	m.ascending = Ascending
	m.mergesort(0, len(m.array)-1)
	return m.array
}
