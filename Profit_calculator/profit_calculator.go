package main

import(
	"fmt"
)

func main() {
	revenue := int(getUserInput("Enter revenue: "))
	expenses := int(getUserInput("Enter expenses: "))
	tax_rate := getUserInput("Enter tax rate: ")

	EBT , profit, ratio := calculateProfit(revenue, expenses, tax_rate)
	fmt.Printf("EBT : %v\nprofit : %v\nratio: %.5f",EBT,profit,ratio)
}

func getUserInput(infoText string) (float64){
	var input float64
	fmt.Print(infoText)
	fmt.Scan(&input)
	return input
}

func calculateProfit(revenue, expenses int, tax_rate float64) (int, float64, float64) {
	EBT := revenue - expenses //earnings_before_tax
	profit := float64(EBT) * (1 - tax_rate/100) //earnings_after_tax
	ratio := float64(EBT) / profit
	return EBT, profit, ratio
}