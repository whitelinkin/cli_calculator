# 🧮 CLI_Calculator 

Простой калькулятор, написанный на Go.  
Поддерживает базовые арифметические операции и хранение истории вычислений.  
Проект создан для практики работы с **Go**, модульной архитектурой и тестированием.

A simple calculator written in Go. 
Supports basic arithmetic operations and calculation history storage. 
The project is created for practicing **Go**, modular architecture, and testing.

---

## 🚀 Возможности | Features
- ➕ Сложение/Addition 
- ➖ Вычитание/Subtraction
- ✖ Умножение/Multiplication
- ➗ Деление (с проверкой деления на ноль)/Division (with zero division check)
- 📜 История вычислений/History of calculations
- 🧪 Unit-тесты для операций и истории/Unit tests for operations and history

---

## 📂 Структура проекта | Project structure
```
calculator/
├── cmd/app/main.go                # Точка входа (CLI)           / Entry point
├── internal/calculator/
│   ├── operations.go              # Арифметические операции     / Arithmetic operations
│   ├── history.go                 # Работа с историей           / Working with history
├── internal/cli/input.go          # Обработка ввода             / Input processing
├── tests/                         # Тесты                       / Unit tests
```
---

## ⚙️ Установка и запуск | Installation and launch
```
Клонировать репозиторий / Clone a repository:
       git clone https://github.com/yourname/calculator.git
       cd calculator/cmd/app
Запустить / To run:
       go run main.go
```
---

## 💡 Использование | Using app
После запуска калькулятор ждёт ввода выражения:
```
Введите выражение: 2 + 2
Результат: 4
```

История хранит прошлые операции:
```
Введите выражение: history
1: 2 + 2 = 4
```

After starting, the calculator waits for you to enter an expression:
```
Enter an expression: 2 + 2
Result: 4
```

The history stores past operations:
```
Enter an expression: history
1: 2 + 2 = 4
```
---

## 🧪 Тестирование | Testing
        Запустить все тесты:
    go test ./...

        Run all tests:

    go test ./...
---

## 📌 Цели проекта
- Изучение структуры проектов на Go  
- Практика с пакетами и модулями  
- Реализовать просто CLI-приложения  
- Научиться писать тесты


   📌 Project goals
- Learning the structure of Go projects 
- Practicing with packages and modules 
- Implementing simple CLI applications 
- Learning how to write tests
