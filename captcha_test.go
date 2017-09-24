package captcha

import (
	"testing"
)

func TestPattern2Put1RightOperandShouldBe1(t *testing.T) {
	actual := right(2, 1, 1, 1)
	expect := "1"
	if actual != expect {
		t.Error("expect", expect, "but", actual)
	}
}

func TestPut1RightOperandShouldBeOne(t *testing.T) {
	actual := right(1, 1, 1, 1)
	expect := "One"
	if actual != expect {
		t.Error("expect", expect, "but", actual)
	}
}
func TestPut5RightOperandShouldBeFive(t *testing.T) {
	actual := right(1, 1, 1, 5)
	expect := "Five"
	if actual != expect {
		t.Error("expect", expect, "but", actual)
	}
}

func TestPut9RightOperandShouldBeNine(t *testing.T) {
	actual := right(1, 1, 1, 9)
	expect := "Nine"
	if actual != expect {
		t.Error("expect", expect, "but", actual)
	}
}
func TestPut1LeftOperandShouldBe1(t *testing.T) {
	actual := left(1, 1)
	expect := "1"
	if actual != expect {
		t.Error("expect", expect, "but", actual)
	}
}

func TestPut5LeftOperandShouldBe5(t *testing.T) {
	actual := left(1, 5)
	expect := "5"
	if actual != expect {
		t.Error("expect", expect, "but", actual)
	}
}

func TestPut9LeftOperandShouldBe9(t *testing.T) {
	actual := left(1, 9)
	expect := "9"
	if actual != expect {
		t.Error("expect", expect, "but", actual)
	}
}

func TestPattern2Put1LeftOperandShouldBeOne(t *testing.T) {
	actual := left(2, 1)
	expect := "One"
	if actual != expect {
		t.Error("expect", expect, "but", actual)
	}
}
func TestPattern2Put5LeftOperandShouldBeFive(t *testing.T) {
	actual := left(2, 5)
	expect := "Five"
	if actual != expect {
		t.Error("expect", expect, "but", actual)
	}
}

func TestPattern2Put9LeftOperandShouldBeNine(t *testing.T) {
	actual := left(2, 9)
	expect := "Nine"
	if actual != expect {
		t.Error("expect", expect, "but", actual)
	}
}
