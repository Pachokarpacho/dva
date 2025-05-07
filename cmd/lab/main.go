package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

// Функция для проверки, является ли число простым
func isPrime(n int) bool {
	if n < 1 {
		return false
	}
	if n == 2 || n == 3 || n == 1 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Чтение содержимого файла input.txt
	data, err := ioutil.ReadFile("/app/input.txt")
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	// Разделение содержимого на строки
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		input := strings.TrimSpace(line)
		if input == "" {
			continue // пропуск пустых строк
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Ошибка: \"%s\" не является целым.\n", input)
			continue
		}

		if num < 1 {
			fmt.Printf("Ошибка: %d должно быть положительным и >= 1.\n", num)
			continue
		}

		if isPrime(num) {
			fmt.Printf("%d - простое число.\n", num)
		} else {
			fmt.Printf("%d - не является простым числом.\n", num)
		}
	}
}
