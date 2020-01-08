package Sorts

func swap(i int, j int, array []float64) {
	x := array[i]
	array[i] = array[j]
	array[j] = x
}


func  intstofloats(Array []int , array []float64) {
	for _, v := range Array {
		array = append(array, float64(v))
	}

}