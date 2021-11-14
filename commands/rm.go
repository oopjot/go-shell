package commands

import (
	"go-shell/entities"
	"go-shell/utils"
	"errors"
)

func Rm(currDir entities.Dir, args ...string) error {
	if len(args) == 0 {
		return errors.New("rm: missing operand")
	}

	for _, path := range args {
		dest, name := utils.GetDest(path)
		dir, err := utils.Unpath(dest, currDir)
		if err != nil {
			return err
		}
		err = dir.Remove(name)
		if err != nil {
			return err
		}
	}
	return nil
}
