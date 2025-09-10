package calculator

import (
	"testing"

	"github.com/whitelinkin/calculator/internal/calculator"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		expr     string
		expected float64
		wantErr  bool
	}{
		{"2 + 3", 5, false},
		{"10 - 4", 6, false},
		{"6 * 7", 42, false},
		{"8 / 2", 4, false},
		{"8 / 0", 0, true},      // ошибка
		{"привет + 5", 0, true}, // ошибка
		{"2 ^ 3", 0, true},      // ошибка
		{"2 +", 0, true},        // ошибка
	}

	for _, tt := range tests {
		got, err := calculator.Evaluate(tt.expr)
		if (err != nil) != tt.wantErr {
			t.Errorf("Анализ(%q) Ошибка = %v, Ожидание ошибки %v", tt.expr, err, tt.wantErr)
			continue
		}
		if !tt.wantErr && got != tt.expected {
			t.Errorf("Анализ(%q) = %v, Ожидание %v", tt.expr, got, tt.expected)
		}
	}
}
