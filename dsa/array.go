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

func selectionSort(arr []int) []int {
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

func bubbleSort(arr []int) []int {
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

func insertionSort(arr []int) []int {
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

func main() {
	// test data for linear and binary search algos
	// var arr = []int{10, 20, 30, 50, 60, 80, 110, 130, 140, 170}
	// println(linearSearch(arr, 20))
	// println(linearSearch(arr, 190))
	// println(binarySearch(arr, 177))
	// println(binarySearch(arr, 20))

	// test data for all sorting algos
	var arr = []int{2, 4, 1, 5, 7, 3}
	print("Original array : ")
	for k := 0; k < len(arr); k++ {
		print(arr[k], " ")
	}

	// test data for selection sort algo
	// println()
	// var sorted_arr = selectionSort(arr)
	// print("Sorted array : ")
	// for k := 0; k < len(arr); k++ {
	// 	print(sorted_arr[k], " ")
	// }

	// test data for bubble sort algo
	// println()
	// var sorted_arr = bubbleSort(arr)
	// print("Sorted array : ")
	// for k := 0; k < len(arr); k++ {
	// 	print(sorted_arr[k], " ")
	// }

	// test data for insertion sort algo
	// println()
	// var sorted_arr = insertionSort(arr)
	// print("Sorted array : ")
	// for k := 0; k < len(arr); k++ {
	// 	print(sorted_arr[k], " ")
	// }
	println()
}
