package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter Tax Rate (as a percentage): ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Printf("Earnings Before Tax: %.2f\n", ebt)
	fmt.Printf("Profit After Tax: %.2f\n", profit)
	fmt.Printf("Profit Ratio: %.2f\n", ratio)
}
