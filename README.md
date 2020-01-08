# Sorts
A package of different sort algorithm in go

## Quick Sort
Using Quick Sort algorith to sort the array bot ascending and descending by giving Ascending parameter false or true.  
You can pass int array or float64 array to QuickSort function.  
This function does NOT affect the array and returns a []float64 slice.  

You can use it like this:   
```go  
    x := []int{3 , 7 , 10 , 6}  
    y := Sorts.QuickSort(true , x)  
```

## Bubble Sort
Using Bubble Sort algorithm to sort the array both ascending and descending by giving Ascending parameter false or true.  
You can pass int array or float64 array to BubbleSort function.  
This function does NOT affect the array and returns a []float64 slice.    

You can use it like this:  
```go
    x := []int{3 , 7 , 10 , 6}
    y := Sorts.BubbleSort(true , x)
```

