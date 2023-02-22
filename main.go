package main

import (
	"basics/basics"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
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
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	return reader.Text()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
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
					break
				default:
					println("Invalid input please try again")
				}
			}
		}
	}
}
