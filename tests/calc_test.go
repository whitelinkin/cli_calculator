package tests

import (
	"testing"

	"github.com/whitelinkin/calculator/internal/calculator"
)

func TestAddition(t *testing.T) {
	result, err := calculator.Evaluate("2 + 3")
	if err != nil {
		t.Errorf("Ожидалось без ошибки, получили: %v", err)
	}
	if result != 5 {
		t.Errorf("Ожидалось 5, получили %v", result)
	}
}

func TestDivisionByZero(t *testing.T) {
	_, err := calculator.Evaluate("10 / 0")
	if err == nil {
		t.Errorf("Ожидалась ошибка деления на 0, но её нет")
	}
}

func TestInvalidInput(t *testing.T) {
	_, err := calculator.Evaluate("abc + 1")
	if err == nil {
		t.Errorf("Ожидалась ошибка при некорректном вводе")
	}
}
