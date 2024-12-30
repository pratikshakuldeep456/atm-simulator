package pkg

import (
	"bufio"
	"fmt"
	"os"
	"pratikshakuldeep456/atm-stimulator/pkg/service"
	"pratikshakuldeep456/atm-stimulator/pkg/store"
	"strings"

	"github.com/urfave/cli/v2"
)

func StartCLI(ctx *cli.Context) error {
	filepath := "balance.json"
	fs := store.GetStore("file_storage", filepath)
	fmt.Println(fs.CheckBalance())

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("available options")
		fmt.Println("1. Credit")
		fmt.Println("2. Debit")
		fmt.Println("3. CheckBalance")
		fmt.Println("4. exit")
		fmt.Println("choose an option b/w 1-3 type 'exit' or 4 for exit")

		input, _ := reader.ReadString('\n')
		fmt.Println("input", input)
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			err := service.CreditAmount(fs, reader)
			if err != nil {
				fmt.Print(err)
			}
		case "2":
			err := (service.DebitAmount(fs, reader))
			if err != nil {
				fmt.Print(err)
			}
		case "3":
			_, err := service.CheckBalance(fs, reader)
			if err != nil {
				fmt.Print(err)
			}
		case "4", "exit":
			fmt.Println("exisiting...")
			return nil
		default:
			fmt.Println("please select a valid option")
		}
	}

}
