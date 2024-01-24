package main

import (
	"fmt"
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

}

func getUserInput(userText string) float64 {
	var userInput float64
	fmt.Print(userText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinantials(revenue, expenses, taxRate float64)(float64,float64,float64){
	ebt := revenue-expenses
	profit := ebt*(1-taxRate/100)
	ratio := ebt/profit

	return ebt,profit,ratio
}