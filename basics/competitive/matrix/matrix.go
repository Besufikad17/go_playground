package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var input = bufio.NewScanner(os.Stdin)

func get_matrices() ([][]int, [][]int, [][]int) {
	println("Enter row and coloumn separating by sapce")
	input.Scan()
	r, err := strconv.Atoi(strings.Split(input.Text(), " ")[0])
	c, err := strconv.Atoi(strings.Split(input.Text(), " ")[1])

	if err != nil {
		log.Fatal(err)
	}

	a := make([][]int, r)
	b := make([][]int, r)
	result := make([][]int, r)

	for i := range a {
		a[i] = make([]int, c)
		b[i] = make([]int, c)
		result[i] = make([]int, c)
	}

	println("Enter element of the first matrix")
	for i := range a {
		for j := range a[0] {
			input.Scan()
			a[i][j], err = strconv.Atoi(input.Text())
		}
	}

	println("Enter element of the second matrix")
	for i := range a {
		for j := range a[0] {
			input.Scan()
			b[i][j], err = strconv.Atoi(input.Text())
		}
	}
	return a, b, result
}

func get_matrices_for_mul() ([][]int, [][]int, [][]int) {
	println("Enter row and coloumn separating by sapce for matrix a ")
	input.Scan()
	r, err := strconv.Atoi(strings.Split(input.Text(), " ")[0])
	c, err := strconv.Atoi(strings.Split(input.Text(), " ")[1])

	println("Enter row and coloumn separating by sapce for matrix b")
	input.Scan()
	r2, err := strconv.Atoi(strings.Split(input.Text(), " ")[0])
	c2, err := strconv.Atoi(strings.Split(input.Text(), " ")[1])

	if err != nil {
		log.Fatal(err)
	}

	a := make([][]int, r)
	b := make([][]int, r2)
	result := make([][]int, r)

	for i := range a {
		a[i] = make([]int, c)
	}

	for i := range b {
		b[i] = make([]int, c2)
	}

	for i := range result {
		result[i] = make([]int, c2)
	}

	println("Enter element of matrix a ")
	for i := range a {
		for j := range a[0] {
			input.Scan()
			a[i][j], err = strconv.Atoi(input.Text())
		}
	}

	println("Enter element of matrix b ")
	for i := range b {
		for j := range b[0] {
			input.Scan()
			b[i][j], err = strconv.Atoi(input.Text())
		}
	}
	return a, b, result

}

func print_menu() string {
	println("\t Welcome to Matrix util")
	println("1. Addition")
	println("2. Substraction")
	println("3. Multiplication")
	println("4. Determinant")
	println("5. Exit")
	input.Scan()
	return input.Text()
}

func main() {
main_loop:
	var choice = print_menu()

	for {
		switch choice {
		case "1":
			var a, b, result = get_matrices()

			for i := range a {
				for j := range a[0] {
					result[i][j] = a[i][j] + b[i][j]
				}
			}

			println("Sum")
			for i := range result {
				fmt.Println(result[i])
			}
			println()
			goto main_loop
		case "2":
			var a, b, result = get_matrices()

			for i := range a {
				for j := range a[0] {
					result[i][j] = a[i][j] - b[i][j]
				}
			}

			println("Difference")
			for i := range result {
				fmt.Println(result[i])
			}
			println()
			goto main_loop
		case "3":
			var a, b, result = get_matrices_for_mul()
			for i := range a {
				for j := range b[0] {
					for k := range b {
						result[i][j] = result[i][j] + a[i][k]*b[k][j]
					}
				}
			}

			println("Product")
			for i := range result {
				fmt.Println(result[i])
			}
			println()
			goto main_loop
		case "4":
			println(choice)
		case "5":
			os.Exit(0)
		default:
			println("Please enter valid choice!!")
		}
	}
}
