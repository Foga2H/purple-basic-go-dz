package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("__Калькулятор валют__")

	for {
		currencyMap := map[string]float64{
			"EUR": 0.9,   // Из USD в EUR
			"RUB": 81.07, // Из USD в RUB
		}

		fromCurrency, amount, toCurrency, err := getUserInput()
		if err != nil {
			fmt.Println(err)
			break
		}

		result, err2 := calculateCurrency(fromCurrency, toCurrency, amount, &currencyMap)
		if err2 != nil {
			fmt.Println(err2)
			break
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
	fromCurrency = strings.ToUpper(fromCurrency)
	fmt.Print("Введите число: ")
	_, err := fmt.Scanf("%f", &amount)
	if err != nil {
		return "", 0.0, "", errors.New("invalid input")
	}
	fmt.Print("Введите целевую валюту: ")
	fmt.Scan(&toCurrency)
	toCurrency = strings.ToUpper(toCurrency)

	isValid := checkUserCurrencyInput(fromCurrency, toCurrency)
	if !isValid {
		return fromCurrency, amount, toCurrency, errors.New("incorrect user currency input")
	}

	return fromCurrency, amount, toCurrency, nil
}

func checkUserCurrencyInput(fromCurrency string, toCurrency string) bool {
	if fromCurrency == toCurrency {
		return false
	}

	if fromCurrency != "USD" && fromCurrency != "EUR" && fromCurrency != "RUB" {
		return false
	}

	if toCurrency != "USD" && toCurrency != "EUR" && toCurrency != "RUB" {
		return false
	}

	return true
}

func calculateCurrency(fromCurrency string, toCurrency string, amount float64, currencyMap *map[string]float64) (float64, error) {
	if fromCurrency == "USD" {
		coefficient, ok := (*currencyMap)[toCurrency]
		if !ok {
			return 0.0, errors.New("incorrect user currency input")
		}

		fmt.Println(coefficient, toCurrency, amount)

		return amount * coefficient, nil
	}

	fromCurrencyCoefficient, ok := (*currencyMap)[fromCurrency]
	if !ok {
		return 0.0, errors.New("incorrect user currency input")
	}

	usdAmount := amount / fromCurrencyCoefficient
	if toCurrency == "USD" {
		return usdAmount, nil
	}

	toCurrencyCoefficient, ok := (*currencyMap)[toCurrency]
	if !ok {
		return 0.0, errors.New("incorrect user currency input")
	}
	return usdAmount * toCurrencyCoefficient, nil
}
