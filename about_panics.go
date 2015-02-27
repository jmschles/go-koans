package go_koans

func divideFourBy(i int) int {
	return 4 / i
}

// not sure what we're trying to do here...

const zero = 0

func aboutPanics() {

	n := divideFourBy(2)
	assert(n == 2) // panics are exceptional errors at runtime
}
