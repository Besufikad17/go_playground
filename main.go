package main

import (
	"basics/basics"
	"basics/basics/dsa"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func to_int_arr(arr []string) []int {
	var int_arr []int
	for i := 0; i < len(arr); i++ {
		ie, err := strconv.Atoi(arr[i])
		int_arr = append(int_arr, ie)
		if err != nil {
			log.Fatal(err)
		}
	}
	return int_arr
}

func get_array() []int {
	println("Enter array of integers :")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sinput := strings.Split(scanner.Text(), " ")
	return to_int_arr(sinput)
}

func print_menu() string {
	println("Welcome to GO_PLAYGROUND")
	println("1. For math functions")
	println("2. For array operations")
	println("3. For stack operations")
	println("4. EXIT")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}

func print_math_menu() string {
	println("1. Area of triangle")
	println("2. Volume of a sphere")
	println("3. Absolute value")
	println("4. Factorial")
	println("5. Get factors")
	println("6. Back to main menu")
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	return reader.Text()
}

func print_array_menu() string {
	println("1. Linear search")
	println("2. Binary search")
	println("3. Selection sort")
	println("4. Insertion sort")
	println("5. Bubble sort")
	println("6. Back to main menu")
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	return reader.Text()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
MENU:
	for {
		var c string = print_menu()
		switch c {
		case "1":
			for {
				var c2 string = print_math_menu()
				switch c2 {
				case "1":
					println("Enter height and width :")
					scanner.Scan()
					in := scanner.Text()
					height, err := strconv.ParseFloat(strings.Split(in, " ")[0], 32)
					width, err := strconv.ParseFloat(strings.Split(in, " ")[1], 32)

					println("Area = ", basics.AREA_OF_TRIANGLE(width, height))
					if err != nil {
						log.Fatal(err)
					}
				case "2":
					println("Enter radius of a sphere :")
					scanner.Scan()
					r, err := strconv.ParseFloat(scanner.Text(), 32)
					println("Volume = ", basics.VOLUME_OF_SPHERE(r))
					if err != nil {
						log.Fatal(err)
					}
				case "3":
					println("Enter x :")
					scanner.Scan()
					n, err := strconv.Atoi(scanner.Text())
					println("|x| = ", basics.ABS(n))
					if err != nil {
						log.Fatal(err)
					}
				case "4":
					println("Enter x :")
					scanner.Scan()
					n, err := strconv.Atoi(scanner.Text())
					println("factorial = ", basics.FACTORIAL(n))
					if err != nil {
						log.Fatal(err)
					}
				case "5":
					println("Enter x :")
					scanner.Scan()
					n, err := strconv.Atoi(scanner.Text())
					basics.GET_FACTORS(n)
					if err != nil {
						log.Fatal(err)
					}
				case "6":
					goto MENU
				default:
					println("Invalid input, please try again!!")
				}
			}
		case "2":
			for {
				var c2 string = print_array_menu()
				switch c2 {
				case "1":
					// accepting new array from user
					arr := get_array()

					println("Enter item u want to search :")
					scanner.Scan()
					n, err := strconv.Atoi(scanner.Text())
					index := dsa.LINEAR_SEARCH(arr, n)
					if index == -1 {
						println("Item not found :(")
					} else {
						println("Item found at index = ", index)
					}
					if err != nil {
						log.Fatal(err)
					}
				case "2":
					// accepting new array from user
					arr := get_array()

					println("Enter item u want to search :")
					scanner.Scan()
					n, err := strconv.Atoi(scanner.Text())
					index := dsa.BINARY_SEARCH(arr, n)
					if index == -1 {
						println("Item not found :(")
					} else {
						println("Item found at index = ", index)
					}
					if err != nil {
						log.Fatal(err)
					}
				case "3":
					// accepting new array from user
					arr := get_array()

					// before sorting algo applied
					print("Original array = ")
					for i := 0; i < len(arr); i++ {
						print(arr[i], " ")
					}
					println()

					sorted_arr := dsa.SELECTION_SORT(arr)

					// after sorting algo applied
					print("Sorted array = ")
					for i := 0; i < len(sorted_arr); i++ {
						print(sorted_arr[i], " ")
					}
					println()
				case "4":
					// accepting new array from user
					arr := get_array()

					// before sorting algo applied
					print("Original array = ")
					for i := 0; i < len(arr); i++ {
						print(arr[i], " ")
					}
					println()

					sorted_arr := dsa.INSERTION_SORT(arr)

					// after sorting algo applied
					print("Sorted array = ")
					for i := 0; i < len(sorted_arr); i++ {
						print(sorted_arr[i], " ")
					}
					println()
				case "5":
					// accepting new array from user
					arr := get_array()

					// before sorting algo applied
					print("Original array = ")
					for i := 0; i < len(arr); i++ {
						print(arr[i], " ")
					}
					println()

					sorted_arr := dsa.BUBBLE_SORT(arr)

					// after sorting algo applied
					print("Sorted array = ")
					for i := 0; i < len(sorted_arr); i++ {
						print(sorted_arr[i], " ")
					}
					println()
				case "6":
					goto MENU
				default:
					println("Invalid input, please try again!!")
				}
			}
		case "3":
			println("case 3")
		case "4":
			os.Exit(0)
		default:
			println("Invalid input!!")
		}
	}
}
