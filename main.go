package main

import "fmt"

func main() {
	transactions := []float64{}
	var a float64
	var sum float64
	for {
		fmt.Print("Введиет свою транзакцию: ")
		fmt.Scan(&a)
		if a == 0 {
			break
		}
		transactions = append(transactions, a)
	}
	for _, value := range transactions {
		sum += value
	}
	fmt.Println(transactions)
	fmt.Println(sum)
}
