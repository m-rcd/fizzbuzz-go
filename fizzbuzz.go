package fizzbuzz_go

import "fmt"

func IsDivisibleBy(number, divisor int) bool {
	return number%divisor == 0
}

func FizzbuzzSays(number int) string {
	if IsDivisibleBy(number, 15) {
		return "fizzbuzz"
	} else if IsDivisibleBy(number, 5) {
		return "buzz"
	} else if IsDivisibleBy(number, 3) {
		return "fizz"
	}
	return fmt.Sprintf("%d", number)
}
