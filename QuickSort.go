package Sorts

var array []int
var ascending bool

func swap(i int, j int) {
	x := array[i]
	array[i] = array[j]
	array[j] = x
}

func Partition(low int, high int) (partition int) {
	partition = low
	great := array[high]
	for i := low; i < high; i++ {
		switch ascending {
		case true:
			if array[i] >= great {
				swap(i, partition)
				partition++
			}
		case false:
			if array[i] <= great {
				swap(i, partition)
				partition++
			}
		}
	}
	swap(partition, high)
	return
}

func Sort(low int, high int) {
	if low < high {
		partition := Partition(low, high)
		Sort(low, partition-1)
		Sort(partition+1, high)
	}
}

//Quick Sort Algorithm
func QuickSort(Ascending bool, Array []int) []int {
	array = Array
	Sort(0, len(array)-1)
	return array
}
