package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	userHeight := 1.8
	var userWeight float64 = 95
	var IMT = userWeight / math.Pow(userHeight, IMTPower)
	fmt.Print(IMT)
}
