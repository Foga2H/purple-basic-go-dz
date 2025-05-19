package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var operations = map[string]func(sum float64, transactionsSlice []float64) float64{
	"AVG": calculateAvg,
	"SUM": calculateSum,
	"MED": calculateMed,
}

func main() {
	fmt.Println("__Basic Calculator__")

	for {
		operation, err := getUserInputOperation()
		if err != nil {
			fmt.Println(err)
			continue
		}

		transactionString := getUserInputTransactions()
		result, err := calculateOperation(operation, transactionString)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Selected operation:", operation)
		fmt.Println("Transaction String:", transactionString)
		fmt.Printf("Result: %.2f", result)
		break
	}
}

func calculateAvg(sum float64, transactionsSlice []float64) float64 {
	return sum / float64(len(transactionsSlice))
}

func calculateSum(sum float64, transactionsSlice []float64) float64 {
	return sum
}

func calculateMed(sum float64, transactionsSlice []float64) float64 {
	result := 0.0
	length := len(transactionsSlice)

	if length == 0 {
		result = 0.0
	} else if length%2 == 0 {
		result = (transactionsSlice[length/2] - transactionsSlice[length/2]) / 2
	} else {
		result = transactionsSlice[length/2]
	}

	return result
}

func calculateOperation(operation string, transactions string) (float64, error) {
	var transactionsSlice []float64
	transactionsSplit := strings.Split(transactions, ",")
	for _, transaction := range transactionsSplit {
		transaction = strings.TrimSpace(transaction)

		if transaction == "" {
			return 0.0, errors.New("empty transaction string")
		}

		transactionFloat, err := strconv.ParseFloat(transaction, 64)
		if err != nil {
			return 0.0, err
		}

		transactionsSlice = append(transactionsSlice, transactionFloat)
	}

	sum := 0.0
	for _, transaction := range transactionsSlice {
		sum += transaction
	}

	result := operations[operation]
	if result != nil {
		return result(sum, transactionsSlice), nil
	}

	return 0.0, errors.New("unknown operation")
}

func getUserInputTransactions() string {
	var transactions string
	fmt.Print("Enter transactions with ',' (example: 1,2,3): ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		transactions = scanner.Text()
	}
	return transactions
}

func getUserInputOperation() (string, error) {
	var operation string
	fmt.Print("Enter your operation (AVG, SUM, MED): ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		operation = scanner.Text()
	}

	switch operation {
	case "AVG":
	case "SUM":
	case "MED":
		break
	default:
		return "", errors.New("invalid operation")
	}

	return operation, nil
}
