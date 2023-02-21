package dsa

var STACK []int
var INDEX int = 0

func PUSH(x int) {
	STACK = append(STACK, x)
	INDEX++
}

func IS_EMPTY() bool {
	if len(STACK) < 0 {
		return true
	} else {
		return false
	}
}

func POP() int {
	var val int
	if IS_EMPTY() {
		println("STACK underflow!!")
	} else {
		val = STACK[INDEX-1]
		INDEX--
	}
	return val
}

func TOP() int {
	var val int
	if IS_EMPTY() {
		println("STACK underflow!!")
	} else {
		val = STACK[INDEX-1]
	}
	return val
}
