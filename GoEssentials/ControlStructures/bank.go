package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64,error) {
	data,err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("Failed to get balance from file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse balance")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64){
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile,[]byte(balanceText),0644)
}


func main() {

	var accountBalance,err  = getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------------------")
	}
	var choice int
	fmt.Println("Welcome to GO Bank")

		for {
			fmt.Println("What do you want to do?")
			fmt.Println("1. Check the balance")
			fmt.Println("2. Deposit money")
			fmt.Println("3. Withdraw money")
			fmt.Println("4. Exit")
			fmt.Print("Choose a number:")			
			fmt.Scan(&choice)

		switch choice{
		case 1:
			fmt.Println("Your balance is:",accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Deposit amount:")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0{
				fmt.Println("Amount must be greater than 0")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("New balance:",accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Print("Withdraw amount:")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0{
				fmt.Println("Amount must be greater than 0")
				continue
			}
			if withdrawAmount > accountBalance{
				fmt.Printf("You cannot withdraw more than %v\n",accountBalance)
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("New balance:",accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Print("Goodbye")
			return
		}
	}

	// for i:=0;i<2;i++ {

	// 	fmt.Println("What do you want to do?")
	// 	fmt.Println("1. Check the balance")
	// 	fmt.Println("2. Deposit money")
	// 	fmt.Println("3. Withdraw money")
	// 	fmt.Println("4. Exit")
	// 	fmt.Print("Choose a number:")
	// 	fmt.Scan(&choice)
	
	// 	if choice == 1{
	// 		fmt.Println("Your balance is:",accountBalance)
	// 	} else if choice == 2 {
	// 		var depositAmount float64
	// 		fmt.Print("Deposit amount:")
	// 		fmt.Scan(&depositAmount)
	// 		if depositAmount <= 0{
	// 			fmt.Println("Amount must be greater than 0")
	// 			return
	// 		}
	// 		accountBalance += depositAmount
	// 		fmt.Println("New balance:",accountBalance)
	// 	} else if choice == 3 {
	// 		var withdrawAmount float64
	// 		fmt.Print("Withdraw amount:")
	// 		fmt.Scan(&withdrawAmount)
	// 		if withdrawAmount <= 0{
	// 			fmt.Println("Amount must be greater than 0")
	// 			return
	// 		}
	// 		if withdrawAmount > accountBalance{
	// 			fmt.Printf("You cannot withdraw more than %v",accountBalance)
	// 			return
	// 		}
	// 		accountBalance -= withdrawAmount
	// 		fmt.Println("New balance:",accountBalance)
	// 	} else {
	// 		fmt.Print("Goodbye")
	// 	}
	// }
}