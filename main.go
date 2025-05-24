package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.25

func main() {
	var years, expectedReturnRate, investmentAmount float64

	//fmt.Print("Enter Investment Amount: ")
	outputText("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("Enter Expected Return Rate: ")
	outputText("Enter Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	//fmt.Print("Enter Years: ")
	outputText("Enter Years: ")
	fmt.Scan(&years)

	futureValues, futureRealValues := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValues, "\n", futureRealValues)

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
	//return
}
