package main

func linearSearch(arr []int, item int) int {
	var index int = -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == item {
			index = i
			break
		}
	}
	return index
}

func binarySearch(arr []int, item int) int {
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

func main() {
	var arr = []int{10, 20, 30, 50, 60, 80, 110, 130, 140, 170}
	println(binarySearch(arr, 177))
	println(binarySearch(arr, 20))
}
