package store

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileStorage struct {
}

func NewFileStorage() *FileStorage {
	return &FileStorage{}
}

func (fs *FileStorage) CheckBalance() (int, error) {
	file, _ := os.ReadFile("balance.json")
	var amount Amount
	err := json.Unmarshal(file, &amount)
	if err != nil {
		fmt.Println(err)
	}
	return amount.Balance, nil
}

func (fs *FileStorage) UpdateAmount(amt int) error {
	file, _ := os.ReadFile("balance.json")
	var amount Amount

	err := json.Unmarshal(file, &amount)
	if err != nil {
		fmt.Println(err)
	}
	amount.Balance = amt
	updateData, err := json.Marshal(amount)
	if err != nil {
		fmt.Println(err)
	}
	os.WriteFile("balance.json", updateData, 0644)
	return nil
}

func (fs *FileStorage) CreditAmount(amt int) error {

	balance, err := fs.CheckBalance()
	if err != nil {
		fmt.Println(err)
	}

	balance += amt
	fmt.Println("current balance", balance)
	err = fs.UpdateAmount(balance)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func (fs *FileStorage) DebitAmount(amt int) error {
	balance, err := fs.CheckBalance()
	if err != nil {
		return fmt.Errorf("err occured while fetching balance", err)
	}
	balance -= amt

	fmt.Println("current balace is", balance)
	err = fs.UpdateAmount(balance)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
