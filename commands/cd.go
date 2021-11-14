package commands

import (
	"go-shell/entities"
	"go-shell/utils"
	"errors"
)

func Cd(currDir entities.Dir, p *entities.Dir, args ...string) error {
	var dest entities.Dir
	var err error
	if len(args) == 0 {
		var root entities.Dir
		root, err = utils.GetRoot(currDir)
		if err != nil {
			return err
		}
		*p = root
		return nil
	}
	if len(args) > 1 {
		return errors.New("cd: too many arguments")
	}
	path := args[0]
	dest, err = utils.Unpath(path, currDir)
	if err != nil {
		return err
	}
	*p = dest
	return nil

}
