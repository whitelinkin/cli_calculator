package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Evaluate(expr string) (float64, error) {
	parts := strings.Fields(expr)
	if len(parts) != 3 {
		return 0, errors.New("Неправильный ввод. Формат: 'Число' 'Оператор' 'Число'")
	}

	a, err1 := strconv.ParseFloat(parts[0], 64)
	b, err2 := strconv.ParseFloat(parts[2], 64)
	if err1 != nil || err2 != nil {
		return 0, errors.New("Неправильные числа")
	}

	operation := parts[1]

	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("На 0 делить нельзя")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("Неизвестный оператор: %s", operation)
	}
}
