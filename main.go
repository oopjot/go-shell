package main

import (
	"go-shell/commands"
	"go-shell/entities"
	"fmt"
	"bufio"
	"os"
)

func main() {
	var input string
	root := entities.RootDir()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = scanner.Text()
		err := commands.Command(root, input)
		if err != nil {
			fmt.Println(err)
		}
	}
}
