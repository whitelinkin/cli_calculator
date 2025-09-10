package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/whitelinkin/calculator/internal/calculator"
)

func Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("'history' - Просмотреть историю.\n'exit' - Выход из программы.\n")
		fmt.Print("Введите выражение(2 + 6): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Выходим...")
			break
		}

		if input == "history" {
			history := calculator.GetHistory()
			if len(history) == 0 {
				fmt.Println("Пусто...")
			} else {
				fmt.Println("История вычислений:")
				for _, entry := range history {
					fmt.Println(" -", entry)
				}
			}
			continue
		}

		result, err := calculator.Evaluate(input)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		calculator.AddHistory(input, result)

		fmt.Println("Результат:", result)
	}
}
