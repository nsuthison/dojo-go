package solutions

import (
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

	carry := 0
	toReturn := ""
	idxNum1 := len(num1) - 1
	idxNum2 := len(num2) - 1

	for idxNum1 >= 0 || idxNum2 >= 0 {

		num1Value := tryGetValueFrom(num1, idxNum1)
		num2Value := tryGetValueFrom(num2, idxNum2)

		sum := num1Value + num2Value + carry
		carry = 0

		if sum > 9 {
			carry = 1
			sum = sum % 10
		}

		toReturn = strconv.Itoa(sum) + toReturn

		idxNum1--
		idxNum2--
	}

	if carry != 0 {
		toReturn = strconv.Itoa(carry) + toReturn
	}

	return toReturn
}

func tryGetValueFrom(number string, index int) int {
	if index < 0 || index >= len(number) {
		return 0
	}

	return mapRuneToInt[rune(number[index])]
}
