package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	operation, numbers := dialogue()
	result := 0.0
	switch operation {
	case "AVG":
		for _, s := range numbers {
			result += s
		}
		result = result / float64(len(numbers))
		fmt.Printf("Результат: %.2f", result)
	case "SUM":
		for _, s := range numbers {
			result += s
		}
		fmt.Printf("Результат: %.2f", result)
	case "MED":
		slices.Sort(numbers)
		if len(numbers)%2 == 0 {
			result = numbers[len(numbers)/2]
			result += numbers[len(numbers)/2-1]
			result /= 2
			fmt.Printf("Результат: %.2f", result)
		} else {
			result = numbers[(len(numbers)+1)/2-1]
			fmt.Printf("Результат: %.2f", result)
		}
	}
}

func dialogue() (string, []float64) {
	var operation string
	var numbers string

	for {
		fmt.Print("Введите тип операции(AVG/SUM/MED): ")
		fmt.Scan(&operation)
		if operation == "AVG" || operation == "SUM" || operation == "MED" {
			break
		}
		fmt.Println("Неправильный ввод.")
	}

	fmt.Scanln()

	fmt.Println("Введите ряд чисел через запятую:")
	fmt.Scanln(&numbers)

	nums := strings.Split(numbers, ",")
	floatNums := make([]float64, 0, len(nums))

	for _, s := range nums {
		trimmed := strings.TrimSpace(s)
		val, err := strconv.ParseFloat(trimmed, 64)
		if err != nil {
			fmt.Printf("%s не число, пропускается.\n", trimmed)
			continue
		}
		floatNums = append(floatNums, val)
	}
	return operation, floatNums
}
