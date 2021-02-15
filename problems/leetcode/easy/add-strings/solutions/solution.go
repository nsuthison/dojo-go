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
	// lenghtToUse := 0

	// if len(num1) < len(num2) {
	// 	lenghtToUse = len(num1)
	// } else {
	// 	lenghtToUse = len(num2)
	// }

	// sum := 0
	// for idx := 0; idx < lenghtToUse; idx++ {

	// }

	return strconv.Itoa(toIntFrom(num1) + toIntFrom(num2))
	// return sum
}

func getValueFrom(number rune, powerFactor float64) int {
	multiplyBy := math.Pow(10, powerFactor)

	return mapRuneToInt[number] * int(multiplyBy)
}

func toIntFrom(number string) int {
	toReturn := 0
	powerFactor := 0.00

	for idx := len(number) - 1; idx >= 0; idx-- {
		letter := rune(number[idx])

		toReturn += getValueFrom(letter, powerFactor)

		powerFactor++
	}

	return toReturn
}
