package main

import "fmt"

func main() {
	const EUR float64 = 0.89
	const RUB float64 = 79.85

	var eurAmount float64 = 10

	usdAmount := eurAmount * EUR
	rubAmount := usdAmount * RUB

	fmt.Printf("RUB: %.2f\n", rubAmount)
}

func readUserInput() string {
	var result string
	fmt.Scan(&result)
	return result
}

func calculateCurrency(amount float64, fromCurrency string, toCurrency string) float64 {

}