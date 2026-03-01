package main

import (
	"fmt"
	"math"
)

func main() {
	userWeight, userHeight := GetUserInput()
	fmt.Print(Result(userWeight, userHeight))
}

func Result(userWeight float64, userHeight float64) string {
	const IMTPower = 2
	var IMT = userWeight / math.Pow(userHeight, IMTPower)
	return fmt.Sprintf("Ваш индекс массы тела: %.1f", IMT)
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
