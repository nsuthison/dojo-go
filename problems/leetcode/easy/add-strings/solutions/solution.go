package solutions

import (
	"math"
	"strconv"
)

var mapRuneToInt = map[rune]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func addStrings(num1 string, num2 string) string {
	return strconv.Itoa(toIntFrom(num1) + toIntFrom(num2))
}

func toIntFrom(number string) int {
	toReturn := 0
	powerFactor := 0.00

	for idx := len(number) - 1; idx >= 0; idx-- {
		letter := rune(number[idx])
		multiplyBy := math.Pow(10, powerFactor)

		toReturn += mapRuneToInt[letter] * int(multiplyBy)

		powerFactor++
	}

	return toReturn
}
