package entities

import (
	"errors"
	"strings"
	"fmt"
)

type file struct {
	name string
	content string
	parent Dir
}

type File interface {
	Name() string
	Read() string
	Write(content string)
	Append(line string)
	ChangeParent(to Dir)
	Rename(name string) error
}

func NewFile(name string, parent Dir) (File, error) {
	newFile := &file{
		name: name,
		content: "",
		parent: parent,
	}
	err := parent.AddFile(newFile)
	return newFile, err
}

func (f *file) Name() string {
	return f.name
}

func (f *file) Read() string {
	return f.content
}

func (f *file) Write(content string) {
	f.content = content + "\n"
}

func (f *file) Append(line string) {
	f.content = f.content + line + "\n"
}

func (f *file) ChangeParent(to Dir) {
	f.parent = to
}

func (f *file) Rename(name string) error {
	if name == f.Name() {
		return nil
	}
	if len(name) == 0 {
		return errors.New("Name cannot be empty.")
	}
	if strings.ContainsRune(name, '/') {
		return errors.New("Name cannot contain '/'.")
	}
	if (f.parent.Exists(name)) {
		msg := fmt.Sprintf("'%s': Already exists", name)
		return errors.New(msg)
	}
	f.name = name
	return nil
}
