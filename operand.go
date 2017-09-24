package captcha

import (
	"strconv"
)

type operand struct {
	pattern       int
	numberOperand int
}

var numbers = map[int]string{1: "One", 2: "Two", 3: "Three", 4: "Four", 5: "Five", 6: "Six", 7: "Seven", 8: "Eight", 9: "Nine"}

func (op operand) left() string {
	if op.pattern == 1 {
		return strconv.Itoa(op.numberOperand)
	}
	return numbers[op.numberOperand]
}

func (op operand) right() string {
	if op.pattern == 2 {
		return numbers[op.numberOperand]
	}
	return strconv.Itoa(op.numberOperand)
}
