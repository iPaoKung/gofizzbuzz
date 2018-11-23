package count

import (
	"strconv"
)

//FizzBuzz counting style
func FizzBuzz(n int) string {
	if numberModByThree(n) && numberModByFive(n) {
		return "FizzBuzz"
	}
	if numberModByThree(n) {
		return "Fizz"
	}
	if numberModByFive(n) {
		return "Buzz"
	}
	return strconv.Itoa(n)
}

func numberModByThree(n int) bool {
	return n%3 == 0
}

func numberModByFive(n int) bool {
	return n%5 == 0
}
