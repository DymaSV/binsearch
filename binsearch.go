package binsearch

import "fmt"

func BinSearch(array *[]int, searchValue int, index int) int {
	arr := *array
	fmt.Println(arr)
	if len(arr) > 1 {
		if arr[len(arr)/2] == searchValue {
			return len(arr) / 2
		} else if arr[len(arr)/2] > searchValue {
			slice1 := arr[0 : len(arr)/2]
			return BinSearch(&slice1, searchValue, index)
		} else if arr[len(arr)/2] < searchValue {
			slice2 := arr[len(arr)/2 : len(arr)-1]
			return index + BinSearch(&slice2, searchValue, len(arr)/2)
		}
	} else {
		if arr[0] == searchValue {
			return 0
		}
	}
	return -1
}
