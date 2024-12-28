package pkg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/gommon/log"
	"github.com/urfave/cli/v2"
)

func StartCLI(ctx *cli.Context) error {

	// READ: how it works
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("available options")
		fmt.Println("1. Credit")
		fmt.Println("2. Debit")
		fmt.Println("3. exit")
		fmt.Println("choose an option b/w 1-2 type 'exit' or 3 for exit")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("enter the amount for credit")

			inputAmount, _ := reader.ReadString('\n')
			inputAmount = strings.TrimSpace(inputAmount)
			creditAmount, _ := strconv.ParseFloat(inputAmount, 64)
			CreditAmount(creditAmount)

		case "2":
			fmt.Println("enter the amount for debit")
			inputAmount, _ := reader.ReadString('\n')
			inputAmount = strings.TrimSpace(inputAmount)
			debitAmount, _ := strconv.ParseFloat(inputAmount, 64)
			DebitAmount(debitAmount)
		case "3":
			fmt.Println("exisiting...")
			return nil
		case "exit":
			fmt.Println("exisiting...")
			return nil
		default:
			fmt.Println("please select a valid option")
		}
	}

	// return nil
}

func CreditAmount(amount float64) error {

	balance := checkBalance()
	balance.Balance = balance.Balance + int(amount)
	updateData, _ := json.Marshal(balance)
	updateDataa := os.WriteFile("balance.json", updateData, 0644)
	fmt.Println(" current balance is", balance)

	if updateDataa != nil {
		log.Printf(" some error occured")
	}
	return nil
}

func DebitAmount(amount float64) error {

	balance := checkBalance()

	if int(amount) > (balance.Balance) {
		log.Error("input amount is more than available balance")
		return fmt.Errorf("input amount is more than available balance")

	}

	balance.Balance = balance.Balance - int(amount)

	updateData, _ := json.Marshal(balance)
	updateDataa := os.WriteFile("balance.json", updateData, 0644)
	if updateDataa != nil {
		log.Printf("input amount is more than available balance")

	}
	fmt.Println(" current balance is", balance)
	return nil

}

type Amount struct {
	Balance int `json:"balance"`
}

func checkBalance() Amount {
	file, _ := os.ReadFile("balance.json")
	var amount Amount
	err := json.Unmarshal(file, &amount)
	if err != nil {
		fmt.Println(err)
	}
	return amount
}
