package main

import (
	"fmt"
)

func merge(arr *[]int, start, mid, end int){
	a := *arr
	p := start
	q := mid + 1

	newArr := make([]int, end-start+1)
	k := 0

	for i := start; i <= end; i++ {
		if p > mid {
			newArr[k] = a[q]
			q += 1
		} else if q > end {
			newArr[k] = a[p]
			p += 1
		} else if a[p] < a[q] {
			newArr[k] = a[p]
			p += 1			
		} else {
			newArr[k] = a[q]
			q += 1						
		}
			k += 1
	}

	for p := 0; p < k; p++ {
		a[start] = newArr[p]
		start += 1
	}

	*arr = a

}

func mergeSort(arr *[]int, start, end int){
	if start < end {
		
		mid := (start + end) / 2

		//fmt.Println(mid)

		mergeSort(arr, start, mid)
		mergeSort(arr, mid+1, end)

		merge(arr, start, mid, end)
	}
}

func main() {
	arr := []int{3,4,5,2,4,7,12,67,1}

	fmt.Println(arr)

	mergeSort(&arr, 0 , len(arr)-1)

	fmt.Println(arr)
}
