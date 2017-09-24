package captcha

import (
	"strconv"
)

var numbers = map[int]string{1: "One", 2: "Two", 3: "Three", 4: "Four", 5: "Five", 6: "Six", 7: "Seven", 8: "Eight", 9: "Nine"}

func left(pattern, leftOperand int) string {
	if pattern == 2 {
		return numbers[leftOperand]
	}
	return strconv.Itoa(leftOperand)
}

func right(pattern, leftOperand, operator, rightOperand int) string {
	if pattern == 2 {
		return strconv.Itoa(rightOperand)
	}
	return numbers[rightOperand]
}
