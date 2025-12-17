package main

import (
	"fmt"
	"example.com/ATM/fileops"
)

const accountBalanceFile = "balance.txt"



func main(){
	accountBalance, err := fileops.ReadBalanceFromFile(accountBalanceFile)
	if err != nil {
		fmt.Printf("Error reading balance: %v\n--------------------------------------------------\n", err)
		panic("Can't continue, sorry!")
	}
	fmt.Println("Welcome to the ATM!")
	var choice int
	fmt.Println("Please choose an option:")
	
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	fmt.Scan(&choice)
	
	switch choice {
	case 1:
		fmt.Printf("Your current balance is: $%.2f\n", accountBalance)
	case 2:
		var depositAmount float64
		fmt.Print("Enter the amount to deposit: ")
		fmt.Scan(&depositAmount)
		if depositAmount <= 0 {
			fmt.Println("Invalid deposit amount!, must be greater than 0")
			return
		}
		accountBalance += depositAmount
		fileops.WriteBalanceToFile(accountBalanceFile, accountBalance)
		fmt.Printf("You have successfully deposited $%.2f. Your new balance is: $%.2f\n", depositAmount, accountBalance)
	case 3:
		var withdrawAmount float64
		fmt.Print("Enter the amount to withdraw: ")
		fmt.Scan(&withdrawAmount)
		if withdrawAmount > accountBalance {
			fmt.Println("Insufficient funds!")
		} else {
			accountBalance -= withdrawAmount
			fileops.WriteBalanceToFile(accountBalanceFile, accountBalance)
			fmt.Printf("You have successfully withdrawn $%.2f. Your new balance is: $%.2f\n", withdrawAmount, accountBalance)
		}
	case 4:
		fmt.Println("Thank you for using the ATM. Goodbye!")
	default:
		fmt.Println("Invalid choice. Please try again.")
	}

	// if choice == 1 {
	// 	fmt.Printf("Your current balance is: $%.2f\n", accountBalance)
	// }
	// if choice == 2 {
	// 	var depositAmount float64
	// 	fmt.Print("Enter the amount to deposit: ")
	// 	fmt.Scan(&depositAmount)
	// 	if depositAmount <= 0 {
	// 		fmt.Println("Invalid deposit amount!, must be greater than 0")
	// 		return
	// 	}
	// 	accountBalance += depositAmount
	// 	fmt.Printf("You have successfully deposited $%.2f. Your new balance is: $%.2f\n", depositAmount, accountBalance)
	// }
	// if choice == 3 {
	// 	var withdrawAmount float64
	// 	fmt.Print("Enter the amount to withdraw: ")
	// 	fmt.Scan(&withdrawAmount)
	// 	if withdrawAmount > accountBalance {
	// 		fmt.Println("Insufficient funds!")
	// 	} else {
	// 		accountBalance -= withdrawAmount
	// 		fmt.Printf("You have successfully withdrawn $%.2f. Your new balance is: $%.2f\n", withdrawAmount, accountBalance)
	// 	}
	// }
	// if choice == 4 {
	// 	fmt.Println("Thank you for using the ATM. Goodbye!")
	// } else {
	// 	fmt.Println("Invalid choice. Please try again.")
	// }
}