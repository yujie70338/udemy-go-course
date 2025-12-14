package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5   // const is immutable
	var investmentAmount = 1000 // var is mutable
	// var expectedReturnRate = 5.5
	expectedReturnRate := 5.5 // shorter way to assign a variable
	var years = 10.0

	fmt.Print("input your investmentAmount : ")
	fmt.Scan(&investmentAmount) // fmt.Scan(&investmentAmount) reads text from standard input (usually the keyboard), converts it to the appropriate type, and stores it into investmentAmount.

	fmt.Print("input your invest year : ")
	fmt.Scan(&years)

	// var futureValue = float64(investmentAmount) * math.Pow((1+expectedReturnRate/100), float64(years))
	futureValue := float64(investmentAmount) * math.Pow((1+expectedReturnRate/100), float64(years))

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
