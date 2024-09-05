package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		investment_amount    float64
		expected_return_rate float64
		years                float64
	)
	fmt.Print("Enter the Investment Amount: ")
	fmt.Scan(&investment_amount)
	fmt.Print("Enter the Return Rate: ")
	fmt.Scan(&expected_return_rate)
	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	future_value := investment_amount * math.Pow(1+expected_return_rate/100, years)
	fmt.Printf("Your current investment will turn to %v in %v years", future_value, years)
}
