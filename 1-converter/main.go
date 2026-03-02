package main

import (
	"fmt"
)

func main() {
	var flag string
	fmt.Printf("Деньги в пересчете: %v\n", calculate(dialogue()))
	for {
		fmt.Println("Продолжить перевод валют?(Yes/No)")
		fmt.Scan(&flag)
		switch flag {
		case "Yes":
			fmt.Printf("Деньги в пересчете: %v\n", calculate(dialogue()))
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
		if origVal == "Rub" || origVal == "Eur" || origVal == "Usd" {
			break
		} else {
			fmt.Println("Неправильный ввод.")
		}
	}
	for {
		fmt.Print("Введите количество денег: ")
		fmt.Scan(&num)
		if num >= 0 {
			break
		} else {
			fmt.Println("Неправильный ввод.")
		}
	}
	for {
		fmt.Print("Введите конечную валюту(Rub/Eur/Usd): ")
		fmt.Scan(&targetVal)
		if targetVal == "Rub" || targetVal == "Eur" || targetVal == "Usd" {
			break
		} else {
			fmt.Println("Неправильный ввод.")
		}
	}
	return num, origVal, targetVal
}
func calculate(num float64, origVal string, targetVal string) float64 {
	const usdToRub = 77.2
	const usdToEur = 0.85
	const eurToRub = usdToRub / usdToEur
	switch {
	case origVal == "Usd" && targetVal == "Rub":
		return num * usdToRub
	case origVal == "Usd" && targetVal == "Eur":
		return num * usdToEur
	case origVal == "Eur" && targetVal == "Rub":
		return num * eurToRub
	case origVal == "Eur" && targetVal == "Usd":
		return num / usdToEur
	case origVal == "Rub" && targetVal == "Usd":
		return num / usdToRub
	case origVal == "Rub" && targetVal == "Eur":
		return num / eurToRub
	default:
		return num
	}
}
