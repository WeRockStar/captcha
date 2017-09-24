package captcha

import "testing"

func TestFirstPatternLeftOperandShouldBe1(t *testing.T) {
	operand := operand{1, 1}
	expected := "1"
	if operand.left() != expected {
		t.Error("Expected", expected, "but", operand.left())
	}
}

func TestFirstPatternLeftOperandShouldBe5(t *testing.T) {
	operand := operand{1, 5}
	expected := "5"
	if operand.left() != expected {
		t.Error("Expected", expected, "but", operand.left())
	}
}

func TestFirstPatternLeftOperandShouldBe9(t *testing.T) {
	operand := operand{1, 9}
	expected := "9"
	if operand.left() != expected {
		t.Error("Expected", expected, "but", operand.left())
	}
}

func TestSecondPatternLeftOperandShouldBeOne(t *testing.T) {
	operand := operand{2, 1}
	expected := "One"
	if operand.left() != expected {
		t.Error("Expected", expected, "but", operand.left())
	}
}

func TestSecondPatternLeftOperandShouldBeFive(t *testing.T) {
	operand := operand{2, 5}
	expected := "Five"
	if operand.left() != expected {
		t.Error("Expected", expected, "but", operand.left())
	}
}

func TestSecondPatternRightOperandShouldBe1(t *testing.T) {
	operand := operand{1, 1}
	expected := "1"
	if operand.right() != expected {
		t.Error("Expected", expected, "but", operand.right())
	}
}

func TestSecondPatternRightOperandShouldBe5(t *testing.T) {
	operand := operand{1, 5}
	expected := "5"
	if operand.right() != expected {
		t.Error("Expected", expected, "but", operand.right())
	}
}
