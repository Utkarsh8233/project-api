package main

import "fmt"

var (
	scan     float64
	revenue  float64
	expenses float64
	tax_rate float64
)

func userInput(text string) float64 {
	fmt.Print(text)
	fmt.Scan(&scan)
	return scan
}
func calculate_ebt_profit_ratio(revenue, expenses, tax_rate float64) (float64, float64, float64) {
	earning_before_tax := revenue - expenses
	profit := earning_before_tax * (1 - tax_rate/100)
	ratio := revenue / profit
	return earning_before_tax, profit, ratio
}

func main() {
	revenue = userInput("Total Revenue: ")
	expenses = userInput("Total Expenses: ")
	tax_rate = userInput("Tax Rate: ")
	earning_before_tax, profit, ratio := calculate_ebt_profit_ratio(revenue, expenses, tax_rate)
	fmt.Printf("Earning before paying tax is: %.2f ,the net profit is: %.2f and the Ratio is: %.2f", earning_before_tax, profit, ratio)
}
