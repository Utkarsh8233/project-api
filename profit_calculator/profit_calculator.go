package main

import "fmt"

func main() {
	var (
		revenue  float64
		expenses float64
		tax_rate float64
	)
	fmt.Print("Total Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Total Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate: ")
	fmt.Scan(&tax_rate)

	earning_before_tax := revenue - expenses
	profit := earning_before_tax * (1 - tax_rate/100)
	ratio := revenue / profit
	fmt.Printf("Earning before paying tax is: %v ,the net profit is: %v and the Ratio is: %v", earning_before_tax, profit, ratio)
}
