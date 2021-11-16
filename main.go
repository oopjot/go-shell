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
		input = scanner.Text()
		if utils.MyContains(input, '/', '\t', ' '
	)
	}
	root := entities.RootDir()
	currDir := root
	for {
		fmt.Printf("user@mushla %s ω ", utils.Path(currDir))
		scanner.Scan()
		input = scanner.Text()
		err := commands.Command(currDir, &currDir, input)
		if err != nil {
			fmt.Println(err)
		}
	}
}
