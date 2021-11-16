package main

import (
	"go-shell/commands"
	"go-shell/entities"
	"go-shell/utils"
	"fmt"
	"bufio"
	"os"
)

func main() {
	var input string
	var login string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("mushla login: ")
		scanner.Scan()
		login = scanner.Text()
		if utils.Contains(login, '/', '\t', ' ', '\n', '.') {
			fmt.Printf("'%s': invalid login\n", login)
		} else {
			break
		}
	}
	root := entities.RootDir()
	currDir := root
	for {
		fmt.Printf("%s@mushla %s ω ", login, utils.Path(currDir))
		scanner.Scan()
		input = scanner.Text()
		err := commands.Command(currDir, &currDir, input)
		if err != nil {
			fmt.Println(err)
		}
	}
}
