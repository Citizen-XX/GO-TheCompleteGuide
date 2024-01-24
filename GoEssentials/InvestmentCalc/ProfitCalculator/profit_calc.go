package main

import (
	"fmt"
	"os"
)

func main() {

	// var revenue, expenses, taxRate float64
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt,profit,ratio := calculateFinantials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n",ebt)
	fmt.Printf("%.1f\n",profit)
	fmt.Printf("%.1f\n",ratio)
	writeData := fmt.Sprintf("EBT %.2f\nProfit %.2f\nRatio %.2f\n",ebt,profit,ratio)
	os.WriteFile("profit_calc.txt",[]byte(writeData),0644)
}

func getUserInput(userText string) float64 {
	var userInput float64
	fmt.Print(userText)
	fmt.Scan(&userInput)
	if userInput <= 0{
		panic("userInput must be positive")
	}
	return userInput
}

func calculateFinantials(revenue, expenses, taxRate float64)(float64,float64,float64){
	ebt := revenue-expenses
	profit := ebt*(1-taxRate/100)
	ratio := ebt/profit

	return ebt,profit,ratio
}