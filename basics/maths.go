package main

import "math"

func area_of_triangle(width float32, height float32) float32 {
	return 0.5 * width * height
}

func volume_of_sphere(radius float64) float64 {
	return math.Round((4 / 3) * (math.Pow(radius, 3)) * math.Pi)
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func factorial(n int) int {
	var fact int = 1
	for n > 0 {
		fact = fact * n
		n = n - 1
	}
	return fact
}

func getFactors(n int) {
	for i := 1; i <= n; i++ {
		if (n % i) == 0 {
			print(i, " ")
		}
	}
}
