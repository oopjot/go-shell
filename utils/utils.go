package utils

import (
	"go-shell/entities"
	"strings"
)

func GetRoot(dir entities.Dir) (entities.Dir, error) {
	if dir.IsRoot() {
		return dir, nil
	} else {
		parent, err := dir.Parent()
		if err != nil {
			return nil, err
		}
		return GetRoot(parent)
	}
}

func Unpath(path string, currDir entities.Dir) (entities.Dir, error) {
	if path[0] == '/' {
		root, err := GetRoot(currDir)
		if err != nil {
			return nil, err
		}
		return Unpath(string(path[1:]), root)
	}
	pathArr := strings.Split(path, "/")

	if len(pathArr) == 1 {
		return currDir.FindDir(pathArr[0])
	}

	nextDir, err := currDir.FindDir(pathArr[0])
	if err != nil {
		return nil, err
	}
	return Unpath(strings.Join(pathArr[1:], "/"), nextDir)
}
