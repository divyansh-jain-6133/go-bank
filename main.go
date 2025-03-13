package main

import (
	"fmt"
	"go_bank/fileops"
)

const accountBalancefile = "balance.tex"

func main() {

	var accountBalance, err = fileops.GetFloatFromFile(accountBalancefile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------")
		// panic("Can't continue, sorry.")
	}
	fmt.Println("Welcome Go bank!! ")

	for {
		presentOptions()

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
			fileops.WriteFloatToFile(accountBalance, accountBalancefile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalancefile)
		case 4:
			fmt.Println("Thank you for using Go Bank! Goodbye!")
			return
		}
		fmt.Println("Your choice is:", choice)

	}
}
