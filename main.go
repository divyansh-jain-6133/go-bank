package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalancefile = "balance.tex"


func writeFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func main() {

	var accountBalance, err = getFloatFromFile(accountBalancefile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------")
		// panic("Can't continue, sorry.")
	}
	fmt.Println("Welcome Go bank!! ")

	for {
		fmt.Println("Select an option:")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("your balance is:", accountBalance)
		case 2:
			fmt.Println("Enter the amount to deposit:")

			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! Updated amount is:", accountBalance)
			writeFloatToFile(accountBalance, accountBalancefile)
		case 3:
			fmt.Println("Enter the amount to withdraw:")

			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invaloid amount. Must be greater than 0.")
				// return
				continue
			}
			if withdrawalAmount > accountBalance {
				fmt.Println("Insufficient balance to withdraw!!!!")
				// return
				continue
			}
			accountBalance -= withdrawalAmount
			fmt.Println("Withdraw successsfull! Updated balance is:", accountBalance)
			writeFloatToFile(accountBalance, accountBalancefile)
		case 4:
			fmt.Println("Thank you for using Go Bank! Goodbye!")
			return
		}
		fmt.Println("Your choice is:", choice)

	}
}