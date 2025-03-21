package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, err := getUserInput("Enter revenue:")
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Enter expenses:")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("Enter taxRate:")
	if err != nil {
		fmt.Println(err)
		return
	}

	EBT, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n", EBT)
	fmt.Printf("Profit:%.1f\n", profit)
	fmt.Printf("Ratio:%.2f\n", ratio)
	storedResults(EBT, profit, ratio)
}
func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Println(infoText)
	fmt.Scanln(&userInput)

	if userInput <= 0 {

		return 0, errors.New("Invalid input. Number must be positive and must be greater than zero.")
	}

	return userInput, nil
}
func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	EBT := revenue - expenses
	profit := EBT * (1 - taxRate/100)
	ratio := EBT / profit

	return EBT, profit, ratio
}

func storedResults(EBT, profit, ratio float64) {
	Result := fmt.Sprintf("EBT: %.1f\nprofit: %.1f\nratio: %.3f\n", EBT, profit, ratio)
	os.WriteFile("Result.text", []byte(Result), 0644)
}
