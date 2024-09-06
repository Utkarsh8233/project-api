package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
	choice   float64
	deposit  float64
	withdraw float64
)

func writeBalanceToFile(balance float64) {
	os.WriteFile("accountBalance.txt", []byte(fmt.Sprint(balance)), 0644)
}
func getBalanceFromFile() (float64, error) {
	balance_in_byte, err := os.ReadFile("accountBalance.txt")
	if err != nil {
		return 1000, errors.New("Failed to read file!, Default balance set to 1000")
	}
	balance_in_string := string(balance_in_byte)
	balance, err := strconv.ParseFloat(balance_in_string, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse value!, Default balance set to 1000")
	}
	return balance, nil
}
func main() {
	accountBalance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	fmt.Println("Welcome to Go Bank!")
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Your current balance is: ", accountBalance)
			continue
		case 2:
			fmt.Print("How much do you want to deposit? : ")
			fmt.Scan(&deposit)
			if deposit >= 0 {
				accountBalance += deposit
				fmt.Println("Balance Updated! Total Account Balance: ", accountBalance)
				writeBalanceToFile(accountBalance)
				continue
			} else {
				fmt.Println("Invalid Value")
				continue
			}
		case 3:
			fmt.Print("How much do you want to withdraw? : ")
			fmt.Scan(&withdraw)
			if withdraw > 0 && withdraw <= accountBalance {
				accountBalance -= withdraw
				fmt.Println("Balance after withdrawal: ", accountBalance)
			} else {
				fmt.Println("Invalid Amount")
				continue
			}
		default:
			fmt.Println("Exiting :)")
		}
		break
	}

	fmt.Print("Thankyou for choosing Go-Bank!")
}
