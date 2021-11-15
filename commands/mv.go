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

	var toMove []interface{}
	last := len(args) - 1

	for _, path := range args[:last] {
		parentPath, name := utils.GetDest(path)
		parent, err := utils.Unpath(parentPath, currDir)
		if err != nil {
			return err
		}
		dir, err := parent.FindDir(name)
		if err != nil {
			file, err := parent.FindFile(name)
			if err != nil {
				return err
			} else {
				toMove = append(toMove, file)
			}
		} else {
				toMove = append(toMove, dir)
			}
	}

	destParentPath, destName := utils.GetDest(args[last])
	destParent, err := utils.Unpath(destParentPath, currDir)
	if err != nil {
		return err
	}
	destDir, err := destParent.FindDir(destName)
	if err != nil {
		destFile, err := destParent.FindFile(destName)
		if err != nil {
			if len(toMove) == 1 {
				switch toMove[0].(type) {
					case entities.File:
						toMove[0].(entities.File).Rename(destName)
						return nil
					case entities.Dir:
						toMove[0].(entities.Dir).Rename(destName)
						return nil
					default:
						return errors.New("mv: unknown error")
				}
			}
			return err
		}

		if len(toMove) > 1 {
			return errors.New(fmt.Sprintf("mv: target '%s' is not a directory", destName))
		}
		switch toMove[0].(type) {
			case entities.Dir:
				return errors.New(fmt.Sprintf("mv: target '%s' is not a directory", destName))
			default:
		}
		return handleFile(destFile, toMove[0].(entities.File))
	}
	return handleDir(destDir, toMove...)
}


func handleDir (dest entities.Dir, source ...interface{}) error {
	for _, entity := range source {
		switch entity.(type){
			case entities.File:
				name := entity.(entities.File).Name()
				_, err := dest.FindDir(name)
				if err == nil {
					return errors.New(fmt.Sprintf("mv: '%s' already exists", name))
				}
				_, err = dest.FindFile(name)
				if err == nil {
					dest.Remove(name)
				}
				entity.(entities.File).Parent().Remove(name)
				entity.(entities.File).ChangeParent(dest)
				err = dest.AddFile(entity.(entities.File))
				if err != nil {
					return err
				}
			case entities.Dir:
				name := entity.(entities.Dir).Name()
				_, err := dest.FindFile(name)
				if err == nil {
					return errors.New(fmt.Sprintf("mv: '%s' already exists", name))
				}
				_, err = dest.FindDir(name)
				if err == nil {
					dest.Remove(name)
				}
				parent, _ := entity.(entities.Dir).Parent()
				parent.Remove(name)
				entity.(entities.Dir).ChangeParent(dest)
				dest.AddDir(entity.(entities.Dir))
			default:
				return errors.New("mv: unknown error")
		}
	}
	return nil
}

func handleFile (dest entities.File, source entities.File) error {
	name := source.Name()
	source.Parent().Remove(name)
	source.ChangeParent(dest.Parent())
	dest.Parent().Remove(name)
	dest.Parent().AddFile(source)
	return nil
}

