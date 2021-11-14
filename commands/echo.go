package commands

import (
	"go-shell/entities"
	"go-shell/utils"
	"fmt"
	"errors"
)

func Echo(currDir entities.Dir, args ...string) error {
	if len(args) == 0 {
		fmt.Println()
		return nil
	}

	var content string
	var pos int
	mode := 0
	for i, str := range args {
		if str == ">" {
			mode = 1
			pos = i + 1
			break
		}
		if str == ">>" {
			mode = 2
			pos = i + 1
			break
		}
		content = content + str + " "
	}
	if mode == 0 {
		fmt.Println(content)
		return nil
	}

	if len(content) != 0 {
		content = content[:len(content) - 1]
	}

	if len(args) < pos {
		return errors.New("syntax error near unexpected token 'newline'")
	}

	rest := args[pos:]
	path := rest[0]
	dir, filename := utils.GetDest(path)
	dest, err := utils.Unpath(dir, currDir)
	if err != nil {
		return err
	}
	file, err := dest.FindFile(filename)
	if err != nil {
		file, err = entities.NewFile(filename, dest)
		if err != nil {
			return err
		}
	}
	if mode == 1 {
		file.Write(content)
	}
	if mode == 2 {
		file.Append(content)
	}
	return nil
}
