package binsearch

//Binary search algorithm
//return index of searched value
//return -1 if not found
func BinSearch(array *[]int, searchValue int, index int) int {
	arr := *array
	if len(arr) > 1 {
		if arr[len(arr)/2] > searchValue {
			slice := arr[0 : len(arr)/2]
			index = index + 0
			inArr := BinSearch(&slice, searchValue, index)
			if inArr == -1 {
				checkSlice(&slice, searchValue, index)
			}
			return inArr
		} else if arr[len(arr)/2] < searchValue {
			slice := arr[len(arr)/2:]
			index = index + len(arr)/2
			inArr := BinSearch(&slice, searchValue, index)
			if inArr == -1 {
				checkSlice(&slice, searchValue, index)
			}
			return inArr
		}
		return index + len(arr)/2
	} else {
		if arr[0] == searchValue {
			return 0
		}
	}
	return -1
}

func checkSlice(array *[]int, searchValue int, index int) int {
	oneByOneValue := getIndexIfEqual(array, searchValue)
	if oneByOneValue == -1 {
		return oneByOneValue
	}
	return index + oneByOneValue
}

func getIndexIfEqual(array *[]int, searchValue int) int {
	for i, v := range *array {
		if v == searchValue {
			return i
		}
	}
	return -1
}
