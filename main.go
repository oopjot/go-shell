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
	root := entities.RootDir()
	scanner := bufio.NewScanner(os.Stdin)
	currDir := root
	for {
		fmt.Printf("user@go-shell %s ω ", utils.Path(currDir))
		scanner.Scan()
		input = scanner.Text()
		err := commands.Command(currDir, &currDir, input)
		if err != nil {
			fmt.Println(err)
		}
	}
}
