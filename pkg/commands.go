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
		fmt.Println("input", input)
		//input = input + "   "
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("enter the amount for credit")
			creditAmount := readInputAmount(reader)
			CreditAmount(creditAmount)

		case "2":
			fmt.Println("enter the amount for debit")
			debitAmount := readInputAmount(reader)
			DebitAmount(debitAmount)
		case "3", "exit":
			fmt.Println("exisiting...")
			return nil
		default:
			fmt.Println("please select a valid option")
		}
	}

	// return nil
}

func SetCurrency(ctx *cli.Context) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter the currency for transaction")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf(err.Error())
	}
	input = strings.TrimSpace(input)

	var curr Amount
	file, _ := os.ReadFile("balance.json")
	json.Unmarshal(file, &curr)
	fmt.Println("currency is", input)
	curr.Currency = input

	abc, _ := json.Marshal(curr)
	fmt.Println(curr)
	os.WriteFile("balance.json", abc, 0644)

	return nil

}
func readInputAmount(reader *bufio.Reader) float64 {

	inputAmount, _ := reader.ReadString('\n')
	inputAmount = strings.TrimSpace(inputAmount)
	creditAmount, _ := strconv.ParseFloat(inputAmount, 64)
	return creditAmount
}

type Amount struct {
	Balance  int    `json:"balance"`
	Currency string `json:"currency"`
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

func CreditAmount(amount float64) error {

	balance := checkBalance()
	balance.Balance += int(amount)

	updateAmount(balance)
	return nil
}

func DebitAmount(amount float64) error {

	balance := checkBalance()

	if int(amount) > (balance.Balance) {
		log.Errorf("input amount ", amount, " is more than available balance ", balance.Balance)
		return fmt.Errorf("input amount is more than available balance")
	}

	balance.Balance -= int(amount)

	updateAmount(balance)
	return nil

}

func updateAmount(balance Amount) (Amount, error) {
	updateData, _ := json.Marshal(balance)
	updateDataa := os.WriteFile("balance.json", updateData, 0644)
	if updateDataa != nil {
		log.Printf("input amount is more than available balance")

	}
	fmt.Println(" current balance is", balance)
	return balance, nil
}
