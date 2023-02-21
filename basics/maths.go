package basics

import "math"

func init() {
	println("Math pacakge loaded :)")
}

func AREA_OF_TRIANGLE(width float32, height float32) float32 {
	return 0.5 * width * height
}

func VOLUME_OF_SPHERE(radius float64) float64 {
	return math.Round((4 / 3) * (math.Pow(radius, 3)) * math.Pi)
}

func ABS(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func FACTORIAL(n int) int {
	var fact int = 1
	for n > 0 {
		fact = fact * n
		n = n - 1
	}
	return fact
}

func GET_FACTORS(n int) {
	for i := 1; i <= n; i++ {
		if (n % i) == 0 {
			print(i, " ")
		}
	}
}
