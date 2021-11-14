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

func GetDest(path string) (string, string) {
	var pathArr []string
	if path[0] == '/' {
		pathArr = strings.Split(path[1:], "/")
		if len(pathArr) == 0 {
			return "/", pathArr[0]
		} else if len(pathArr) == 1 {
			return ".", pathArr[0]
		} else {
			last := len(pathArr) - 1
			return strings.Join(pathArr[:last], "/"), pathArr[last]
		}
	}
	pathArr = strings.Split(path, "/")
	if len(pathArr) == 1 {
		return ".", pathArr[0]
	} else {
		last := len(pathArr) - 1
		return strings.Join(pathArr[:last], "/"), pathArr[last]
	}
}

func Path(dir entities.Dir) string {
	if dir.IsRoot() {
		return "/"
	}
	parent, _ := dir.Parent()
	if parent.IsRoot() {
		return Path(parent) + dir.Name()
	}
	return Path(parent) + "/" + dir.Name()
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
