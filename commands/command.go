package commands

import (
	"go-shell/entities"
	"fmt"
	"strings"
	"errors"
)

func Command(currDir entities.Dir, input string) error {
	if (len(input)) == 0 {
		return nil
	}

	inputArr := strings.Fields(input)
	args := inputArr[1:]
	switch command := inputArr[0]; command {
		case "ls":
			return Ls(currDir, args...)
		case "touch":
			return Touch(currDir, args...)
		case "mkdir":
			return Mkdir(currDir, args...)
		case "cd":
			fmt.Printf("Got cd in %s\n", currDir.Name())
		case "echo":
			fmt.Printf("Got echo in %s\n", currDir.Name())
		case "cat":
			fmt.Printf("Got cat in %s\n", currDir.Name())
		case "rm":
			fmt.Print("Got rm\n")
		default:
			return errors.New(fmt.Sprintf("'%s': command not found", command))
	}

	return nil
}
