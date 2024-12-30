package store

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileStorage struct {
	FilePath string
}

func NewFileStorage(fp string) *FileStorage {
	return &FileStorage{FilePath: fp}
}

func ReadFile(fp string) ([]byte, error) {
	file, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println(err)
	}
	return file, nil
}
func (fs *FileStorage) CheckBalance() (int, error) {
	file, _ := ReadFile(fs.FilePath)
	var amount Amount
	err := json.Unmarshal(file, &amount)
	if err != nil {
		fmt.Println(err)
	}
	return amount.Balance, nil
}

func (fs *FileStorage) UpdateAmount(amt int) error {
	file, _ := ReadFile(fs.FilePath)
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
	os.WriteFile(fs.FilePath, updateData, 0644)
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

	if amt > balance {
		return fmt.Errorf(" input amount is more than available")

	}
	balance -= amt

	fmt.Println("current balace is", balance)
	err = fs.UpdateAmount(balance)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
