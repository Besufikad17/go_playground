package main

import (
	"basics/basics"
	"bufio"
	"fmt"
	"os"
)

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
	reader := bufio.NewReader(os.Stdin)
	choice, err := reader.ReadString('\n')
	if err != nil {
		println("Err occured while reading input")
		return ""
	}
	return choice
}

func main() {
	var c string = print_menu()
	switch c {
	case "1":
		var c2 string = print_math_menu()
		switch c2 {
		case "1":
			var height float32
			var width float32
			println("Enter height and width :")
			fmt.Scanf("%f", &height)
			fmt.Scanf("%f", &width)
			println("Area = ", basics.AREA_OF_TRIANGLE(width, height))
		case "2":
			var radius float32
			println("Enter radius of a sphere :")
			fmt.Scanf("%f", radius)
			println("Volume = ", basics.VOLUME_OF_SPHERE(float64(radius)))
		case "3":
			var n int
			println("Enter x :")
			fmt.Scanf("%i", n)
			println("|x| = ", basics.ABS(n))
		case "4":
			var n int
			println("Enter x :")
			fmt.Scanf("%i", n)
			println("factorial = ", basics.FACTORIAL(n))
		case "5":
			var n int
			println("Enter x :")
			fmt.Scanf("%i", n)
			basics.GET_FACTORS(n)
		case "6":
			// exit
		default:
			println("Invalid input please try again")
		}
	}
}
