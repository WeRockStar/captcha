package captcha

import "testing"

func TestPut1ShoulBePlus(t *testing.T) {
	op := operator{1}
	expected := "+"

	if op.get() != expected {
		t.Error("expected", expected, "but", op.get())
	}
}
func TestPut2ShoulBeMinus(t *testing.T) {
	op := operator{2}
	expected := "-"

	if op.get() != expected {
		t.Error("expected", expected, "but", op.get())
	}
}

func TestPut3ShouldBeMul(t *testing.T) {
	op := operator{3}
	expected := "*"

	if op.get() != expected {
		t.Error("expected", expected, "but", op.get())
	}
}
