package commands

import (
	"go-shell/entities"
	"go-shell/utils"
	"errors"
	"fmt"
)

func Mv(currDir entities.Dir, args ...string) error {
	if len(args) == 0 {
		return errors.New("mv: missing file operand")
	}
	if len(args) == 1 {
		return errors.New(fmt.Sprintf("mv: missing destination operand after '%s'", args[0]))
	}

	last := len(args) - 1

	destParentPath, destName := utils.GetDest(args[last])
	destParentDir, err := utils.Unpath(destParentPath, currDir)
	if err != nil {
		return err
	}

	for i, objPath := range args[:last] {
		objParentPath, objName := utils.GetDest(objPath)
		objParentDir, err := utils.Unpath(objParentPath, currDir)
		if err != nil {
			return err
		}
		destDir, err := destParentDir.FindDir(destName)
		if err != nil {
			if i != 0 {
				return err
			}
			file, err := objParentDir.FindFile(objName)
			if err != nil {
				return err
			}
			destParentDir.Remove(objName)
			file.Rename(objName)
			destParentDir.AddFile(file)
			file.ChangeParent(destParentDir)
		}

		file, nil := objParentDir.FindFile(objName)
		if err != nil {
			// nie znalazlem pliku
			dir, err := objParentDir.FindDir(objName)
			if err != nil {
				return err
			}
			// znalazlem dir
			err = destDir.AddDir(dir)
			if err != nil {
				return err
			}
			dir.ChangeParent(destDir)
		}
		err = destDir.AddFile(file)
		if err != nil {
			destDir.Remove(file.Name())
			destDir.AddFile(file)
		}
		objParentDir.Remove(objName)
		file.ChangeParent(destDir)
	}
	return nil
}
