package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("__Калькулятор валют__")

	for {
		fromCurrency, amount, toCurrency, err := getUserInput()
		if err != nil {
			fmt.Println(err)
			continue
		}

		result, err := calculateCurrency(fromCurrency, toCurrency, amount)
		if err != nil {
			fmt.Println(err)
			continue
		}

		outputResult(fromCurrency, toCurrency, result)
		break
	}
}

func outputResult(fromCurrency string, toCurrency string, result float64) {
	fmt.Println("Результаты: ")
	fmt.Printf("%s -> %s = %.2f\n", fromCurrency, toCurrency, result)
}

func getUserInput() (string, float64, string, error) {
	var fromCurrency, toCurrency string
	var amount float64

	fmt.Print("Введите исходную валюту (USD, EUR, RUB): ")
	fmt.Scan(&fromCurrency)
	validFromCurrency, err := checkUserCurrency(strings.ToUpper(fromCurrency))
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
	validToCurrency, err := checkUserCurrency(strings.ToUpper(toCurrency))
	if err != nil {
		return "", 0, "", err
	}

	if validFromCurrency == validToCurrency {
		return "", 0, "", errors.New("валюты не должны совпадать")
	}

	return strings.ToUpper(validFromCurrency), validAmount, strings.ToUpper(validToCurrency), nil
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

func calculateCurrency(fromCurrency string, toCurrency string, amount float64) (float64, error) {
	currencyMap := map[string]float64{
		"EUR": 0.9,   // Из USD в EUR
		"RUB": 81.07, // Из USD в RUB
	}
	var mapIndex string

	if toCurrency == "USD" {
		mapIndex = fromCurrency
	} else {
		mapIndex = toCurrency
	}

	coefficient, ok := currencyMap[mapIndex]
	fmt.Println(currencyMap, fromCurrency, toCurrency)
	if !ok {
		return 0.0, errors.New("не найдена валюта")
	}

	if toCurrency == "USD" {
		return amount / coefficient, nil
	}

	usdAmount := amount

	if fromCurrency != "USD" {
		usdAmount = amount / coefficient
	}

	return usdAmount * coefficient, nil
}
