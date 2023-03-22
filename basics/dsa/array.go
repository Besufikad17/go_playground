package dsa

func LINEAR_SEARCH(arr []int, item int) int {
	var index int = -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == item {
			index = i
			break
		}
	}
	return index
}

func BINARY_SEARCH(arr []int, item int) int {
	var start int = 0
	var end int = len(arr)
	for start < end {
		var mid int = int((start + (end - 1)) / 2)
		if arr[mid] == item {
			return mid
		} else if arr[mid] > item {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

func SELECTION_SORT(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				var temp int = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

func BUBBLE_SORT(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				var temp int = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}

func INSERTION_SORT(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			var temp int = arr[j-1]
			arr[j-1] = arr[j]
			arr[j] = temp
			j--
		}
	}
	return arr
}
