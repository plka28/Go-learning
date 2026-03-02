package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover", r)
		}
	}()
	i := "Yes"
	for i != "No" {
		userWeight, userHeight := GetUserInput()
		IMT, err := calculateResult(userWeight, userHeight)
		if err != nil {
			// fmt.Println("Ошибка в вводе данных")
			// continue
			panic("Ошибка в вводе данных")
		}
		outputResult(IMT)
		fmt.Println("Продолжить рассчет индекса массы тела?(Yes/No)")
		fmt.Scan(&i)
		for i != "Yes" && i != "No" {
			fmt.Println("Неверный ввод.")
			fmt.Println("Продолжить рассчет индекса массы тела?(Yes/No)")
			fmt.Scan(&i)
		}
	}
}
func calculateResult(userWeight, userHeight float64) (float64, error) {
	const IMTPower = 2
	if userHeight <= 0 || userWeight <= 0 {
		return 0, errors.New("INCORRECT_PARAMS")
	}
	return userWeight / math.Pow(userHeight, IMTPower), nil
}
func outputResult(IMT float64) {
	fmt.Printf("Ваш индекс массы тела: %.1f\n", IMT)
	switch {
	case IMT < 16:
		fmt.Println("Сильный недостаток веса")
	case IMT < 18.5:
		fmt.Println("Задохлик")
	case IMT < 25:
		fmt.Println("Норма")
	case IMT < 30:
		fmt.Println("На грани")
	default:
		fmt.Println("Ожирение")
	}
}

func GetUserInput() (float64, float64) {
	var userHeight float64
	var userWeight float64
	fmt.Print("Введите свой рост(м): ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес(кг): ")
	fmt.Scan(&userWeight)
	return userWeight, userHeight
}
