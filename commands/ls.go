package commands

import (
	"go-shell/entities"
	"go-shell/utils"
	"fmt"
)


func Ls(currDir entities.Dir, args ...string) error {
	if len(args) == 0 {
		fmt.Print(currDir.List())
		return nil
	}
	path := args[0]
	found, err := utils.Unpath(path, currDir)
	if err != nil {
		return err
	}
	fmt.Print(found.List())
	return nil

}
