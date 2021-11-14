package commands

import (
	"go-shell/entities"
	"go-shell/utils"
	"errors"
)

func Mkdir(currDir entities.Dir, args ...string) error {
	if len(args) == 0 {
		return errors.New("mkdir: missing operand")
	}

	for _, path := range args {
		var dest entities.Dir
		var err error

		dirPath, dirname := utils.GetDest(path)
		dest, err = utils.Unpath(dirPath, currDir)
		if err != nil {
			return err
		}
		_, err = entities.NewDir(dirname, dest)
		if err != nil {
			return err
		}
	}

	return nil
}
