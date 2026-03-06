package main

import (
	"fmt"
)

var rates = map[string]float64{
	"Rub": 1,
	"Usd": 77.2,
	"Eur": 77.2 / 0.85,
}

func main() {
	var flag string
	fmt.Printf("Деньги в пересчете: %.2f\n", calculate(dialogue()))

	for {
		fmt.Println("Продолжить перевод валют?(Yes/No)")
		fmt.Scan(&flag)

		switch flag {
		case "Yes":
			fmt.Printf("Деньги в пересчете: %.2f\n", calculate(dialogue()))
		case "No":
			return
		default:
			fmt.Println("Неправильный ввод")
		}
	}
}

func dialogue() (float64, string, string) {
	var origVal string
	var targetVal string
	var num float64

	for {
		fmt.Print("Введите изначальную валюту(Rub/Eur/Usd): ")
		fmt.Scan(&origVal)

		if _, ok := rates[origVal]; ok {
			break
		}
		fmt.Println("Неправильный ввод.")
	}

	for {
		fmt.Print("Введите количество денег: ")
		fmt.Scan(&num)

		if num >= 0 {
			break
		}
		fmt.Println("Неправильный ввод.")
	}

	for {
		fmt.Print("Введите конечную валюту(Rub/Eur/Usd): ")
		fmt.Scan(&targetVal)

		if _, ok := rates[targetVal]; ok {
			break
		}
		fmt.Println("Неправильный ввод.")
	}

	return num, origVal, targetVal
}

func calculate(num float64, origVal string, targetVal string) float64 {
	// сначала переводим исходную сумму в рубли
	inRub := num * rates[origVal]

	// потом из рублей в целевую валюту
	return inRub / rates[targetVal]
}
