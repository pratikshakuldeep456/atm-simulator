package service

import (
	"bufio"
	"fmt"
	"pratikshakuldeep456/atm-stimulator/pkg/store"
	"strconv"
	"strings"
)

func CreditAmount(store store.StoreSerivce, reader *bufio.Reader) error {
	fmt.Println("enter the amount for credit")
	creditAmount := readInputAmount(reader)
	err := store.CreditAmount(int(creditAmount))
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
func DebitAmount(store store.StoreSerivce, reader *bufio.Reader) error {
	fmt.Println("enter the amount for debit")
	debitAmount := readInputAmount(reader)
	err := store.DebitAmount(int(debitAmount))
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func CheckBalance(store store.StoreSerivce, reader *bufio.Reader) (int, error) {
	data, err := store.CheckBalance()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("available balance", data)
	return data, nil
}
func readInputAmount(reader *bufio.Reader) float64 {
	inputAmount, _ := reader.ReadString('\n')
	inputAmount = strings.TrimSpace(inputAmount)
	creditAmount, _ := strconv.ParseFloat(inputAmount, 64)
	return creditAmount
}
