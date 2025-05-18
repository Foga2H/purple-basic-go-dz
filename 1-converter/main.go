package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("__Калькулятор валют__")

	for {
		var fromCurrency string
		var toCurrency string
		var amount float64
		var err error

		currencyMap := map[string]float64{
			"EUR": 0.9,   // Из USD в EUR
			"RUB": 81.07, // Из USD в RUB
		}

		fromCurrency, amount, toCurrency, err = getUserInput()
		if err != nil {
			fmt.Println(err)
			continue
		}

		result, error2 := calculateCurrency(fromCurrency, toCurrency, amount, &currencyMap)
		if error2 != nil {
			fmt.Println(error2)
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
	var validAmount float64

	fmt.Print("Введите исходную валюту (USD, EUR, RUB): ")
	fmt.Scan(&fromCurrency)
	validFromCurrency := strings.ToUpper(fromCurrency)
	fmt.Print("Введите число: ")
	fmt.Scan(&amount)
	checkAmount, err := checkUserAmount(amount)
	if err != nil {
		return "", 0, "", err
	}
	validAmount = checkAmount
	fmt.Print("Введите целевую валюту: ")
	fmt.Scan(&toCurrency)
	validToCurrency := strings.ToUpper(toCurrency)

	err = checkUserCurrencies(validFromCurrency, validToCurrency)
	if err != nil {
		return "", 0, "", err
	}

	return strings.ToUpper(validFromCurrency), validAmount, strings.ToUpper(validToCurrency), nil
}

func checkUserCurrencies(fromCurrency string, toCurrency string) error {
	if toCurrency != "EUR" && toCurrency != "RUB" && toCurrency != "USD" {
		return errors.New("некорректная целевая валюта")
	}

	if fromCurrency != "EUR" && fromCurrency != "RUB" && fromCurrency != "USD" {
		return errors.New("некорректная исходная валюта")
	}

	if toCurrency == fromCurrency {
		return errors.New("валюты не должны совпадать")
	}

	return nil
}

func checkUserAmount(amount float64) (float64, error) {
	if amount <= 0 {
		return 0, errors.New("некорректное число")
	}
	return amount, nil
}

func calculateCurrency(fromCurrency string, toCurrency string, amount float64, currencyMap *map[string]float64) (float64, error) {
	var mapIndex string
	var usdAmount float64

	if toCurrency == "USD" {
		mapIndex = fromCurrency
	} else {
		mapIndex = toCurrency
	}

	currencyMapValue := *currencyMap
	coefficient, ok := currencyMapValue[mapIndex]
	fmt.Println(currencyMapValue, fromCurrency, toCurrency)
	if !ok {
		return 0.0, errors.New("не найдена валюта")
	}

	if toCurrency == "USD" {
		return amount / coefficient, nil
	}

	if fromCurrency != "USD" {
		usdAmount = amount * coefficient
		coefficient = currencyMapValue[fromCurrency]
	} else {
		usdAmount = amount
	}

	return usdAmount * coefficient, nil
}
