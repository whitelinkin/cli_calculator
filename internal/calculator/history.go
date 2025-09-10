package calculator

import "fmt"

var history []string

func AddHistory(expr string, result float64) {
	entry := fmt.Sprintf("%s = %v", expr, result)
	history = append(history, entry)
}

func GetHistory() []string {
	return history
}

func ClearHistory() {
	history = []string{}
}
