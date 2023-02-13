package main

import "fmt"

func main() {
	var total float32 = 0
	var avg float32 = 0
	var number_subjects int = 5
	var count int = 1
	for number_subjects > 0 {
		println("enter subject", count, ":")
		var grade float32
		fmt.Scanln(&grade)
		total = total + grade
		number_subjects--
		count++
	}
	avg = total / 5
	println("total = ", int(total))
	println("average = ", int(avg))
}
