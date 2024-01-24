package main

import (
	"fmt"
	"math"
)

const inflationRate = 4.5

func main(){

	var investmentAmount float64
	var years float64 = 10
	expectedReturnRate := 16.1

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount,expectedReturnRate,years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

func outputText(text string){
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64)(float64,float64){
	fv := investmentAmount*math.Pow(1+expectedReturnRate/100,years)
	rfv := fv/math.Pow(1+inflationRate/100,years)
	return fv,rfv
}