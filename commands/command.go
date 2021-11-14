package commands

import (
	"go-shell/entities"
	"fmt"
	"strings"
	"errors"
)

func Command(currDir entities.Dir, p *entities.Dir, input string) error {
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
			return Cd(currDir, p, args...)
		case "echo":
			return Echo(currDir, args...)
		case "cat":
			return Cat(currDir, args...)
		case "rm":
			return Rm(currDir, args...)
		case "mv":
			return Mv(currDir, args...)
		case "clear":
			return Clear()
		default:
			return errors.New(fmt.Sprintf("'%s': command not found", command))
	}

	return nil
}
