package entities

import (
	"errors"
	"fmt"
)

type rootDir struct {
	files []File
	dirs []Dir
}

func RootDir() Dir {
	return &rootDir{}
}

func (d *rootDir) Name() string {
	return "/"
}

func (d *rootDir) Exists(name string) bool {
	for _, _d := range d.dirs {
		if _d.Name() == name {
			return true
		}
	}
	for _, f := range d.files {
		if f.Name() == name {
			return true
		}
	}
	return false
}


func (d *rootDir) AddFile(newFile File) error {
	if d.Exists(newFile.Name()) {
		msg := fmt.Sprintf("'%s': Already exists", newFile.Name())
		return errors.New(msg)
	}
	d.files = append(d.files, newFile)
	return nil
}

func (d *rootDir) AddDir(newDir Dir) error {
	if d.Exists(newDir.Name()) {
		msg := fmt.Sprintf("'%s': Already extsts", newDir.Name())
		return errors.New(msg)
	}
	d.dirs = append(d.dirs, newDir)
	return nil
}

func (d *rootDir) Remove(name string) error {
	for i, f := range d.files {
		if f.Name() == name {
			bp := len(d.files) - 1
			d.files[i] = d.files[bp]
			d.files = d.files[:bp]
			return nil
		}
	}
	for i, _d := range d.dirs {
		if _d.Name() == name {
			bp := len(d.dirs) - 1
			d.dirs[i] = d.dirs[bp]
			d.dirs = d.dirs[:bp]
			return nil
		}
	}

	msg := fmt.Sprintf("'%s': No such file or directory", name)
	return errors.New(msg)
}

func (d *rootDir) ChangeParent(to Dir) {}

func (d *rootDir) Rename(name string) error {
	return errors.New("You cannot change root name")
}

func (d *rootDir) List() string {
	result := ""
	for _, _d := range d.dirs {
		result = result + "[d] " + _d.Name() + string("\n")
	}
	for _, f := range d.files {
		result = result + "[f] " + f.Name() + string("\n")
	}
	return result
}

func (d *rootDir) IsRoot() bool {
	return true
}
