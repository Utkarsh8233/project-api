package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
	scan            float64
	error_statement error
)

func userInput(text string) (float64, error) {
	fmt.Print(text)
	fmt.Scan(&scan)
	if scan <= 0 {
		err := errors.New("Invalid Value")
		return scan, err
	}
	return scan, nil
}
func calculate_ebt_profit_ratio(revenue, expenses, tax_rate float64) error {
	ebt := revenue - expenses
	profit := ebt * (1 - tax_rate/100)
	ratio := revenue / profit
	err := os.WriteFile("profit.txt", []byte(fmt.Sprint(profit)), 0644)
	if err != nil {
		error_statement = errors.New("Unable to write to file profit.txt")
	}
	err = os.WriteFile("ratio.txt", []byte(fmt.Sprint(ratio)), 0644)
	if err != nil {
		error_statement = errors.New("Unable to write to file ratio.txt")
	}
	err = os.WriteFile("ebt.txt", []byte(fmt.Sprint(ebt)), 0644)
	if err != nil {
		error_statement = errors.New("Unable to write to file ebt.txt")
	}
	return error_statement
}
func get_ebt_profit_ratio() (float64, float64, float64, error) {
	profit, err := os.ReadFile("profit.txt")
	if err != nil {
		error_statement = errors.New("unable to read profit.txt")
	}
	profit_float64, err := strconv.ParseFloat(string(profit), 64)
	if err != nil {
		error_statement = errors.New("unable to parse profit value")
	}
	ratio, err := os.ReadFile("ratio.txt")
	if err != nil {
		error_statement = errors.New("Unable to read ratio.txt")
	}
	ratio_float64, err := strconv.ParseFloat(string(ratio), 64)
	if err != nil {
		error_statement = errors.New("Unable to parse ratio value")
	}
	ebt, err := os.ReadFile("ebt.txt")
	if err != nil {
		error_statement = errors.New("Unable to read ebt.txt")
	}
	ebt_float64, err := strconv.ParseFloat(string(ebt), 64)
	if err != nil {
		error_statement = errors.New("Unable to parse ebt value")
	}
	return profit_float64, ratio_float64, ebt_float64, error_statement

}

func main() {
	revenue, err := userInput("Total Revenue: ")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	expenses, err := userInput("Total Expenses: ")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	tax_rate, err := userInput("Tax Rate: ")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	err = calculate_ebt_profit_ratio(revenue, expenses, tax_rate)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	profit, ratio, ebt, err := get_ebt_profit_ratio()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Printf("Earning before paying tax is: %.2f ,the net profit is: %.2f and the Ratio is: %.2f", ebt, profit, ratio)
}
