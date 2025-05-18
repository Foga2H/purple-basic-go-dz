package main

import (
	"errors"
	"fmt"
)

const EUR float64 = 0.9
const RUB float64 = 81.07

func main() {
	fmt.Println("__Калькулятор валют__")

	for {
		fromCurrency, amount, toCurrency, err := getUserInput()
		if err != nil {
			fmt.Println(err)
			continue
		}

		result := calculateCurrency(fromCurrency, toCurrency, amount)
		outputResult(fromCurrency, toCurrency, result)
		break
	}
}

func outputResult(fromCurrency string, toCurrency string, result float64) {
	fmt.Println("Результаты: ")
	fmt.Printf("%s -> %s = %f\n", fromCurrency, toCurrency, result)
}

func getUserInput() (string, float64, string, error) {
	var fromCurrency, toCurrency string
	var amount float64

	fmt.Print("Введите исходную валюту (USD, EUR, RUB): ")
	fmt.Scan(&fromCurrency)
	validFromCurrency, err := checkUserCurrency(fromCurrency)
	if err != nil {
		return "", 0, "", err
	}
	fmt.Print("Введите число: ")
	fmt.Scan(&amount)
	validAmount, err := checkUserAmount(amount)
	if err != nil {
		return "", 0, "", err
	}
	fmt.Print("Введите целевую валюту: ")
	fmt.Scan(&toCurrency)
	validToCurrency, err := checkUserCurrency(toCurrency)
	if err != nil {
		return "", 0, "", err
	}

	if validFromCurrency == validToCurrency {
		return "", 0, "", errors.New("валюты не должны совпадать")
	}

	return validFromCurrency, validAmount, validToCurrency, nil
}

func checkUserCurrency(currency string) (string, error) {
	if currency != "EUR" && currency != "RUB" && currency != "USD" {
		return "", errors.New("некорректная валюта")
	}
	return currency, nil
}

func checkUserAmount(amount float64) (float64, error) {
	if amount <= 0 {
		return 0, errors.New("некорректное число")
	}
	return amount, nil
}

func calculateCurrency(fromCurrency string, toCurrency string, amount float64) float64 {
	var result float64

	switch fromCurrency {
	case "USD":
		if toCurrency == "EUR" {
			result = amount * EUR
			break
		}

		if toCurrency == "RUB" {
			result = amount * RUB
			break
		}
	case "EUR":
		usdAmount := amount / EUR

		if toCurrency == "USD" {
			result = usdAmount
			break
		}

		if toCurrency == "RUB" {
			result = usdAmount * RUB
			break
		}
	case "RUB":
		usdAmount := amount / RUB

		if toCurrency == "USD" {
			result = usdAmount
			break
		}

		if toCurrency == "EUR" {
			result = usdAmount * EUR
			break
		}
	}

	return result
}
