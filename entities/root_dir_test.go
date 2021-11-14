package entities_test

import (
	"go-shell/entities"
	"testing"
	"fmt"
)

func TestRootDir(t *testing.T) {
	root := entities.RootDir()

	// Assert that root dir name is "/"
	if root.Name() != "/" {
		t.Errorf("Root's name incorrect. Should be '/', got '%s'.", root.Name())
	}

	// Test isRoot()
	if !root.IsRoot() {
		t.Error("Root is root, but got false")
	}

	// Test root's Exists() method.
	dir1Name := "dir1"
	if root.Exists(dir1Name) {
		t.Errorf("'%s' doesn't exist in root, but got true.", dir1Name)
	}
	entities.NewDir(dir1Name, root)
	if !root.Exists(dir1Name) {
		t.Errorf("'%s' exists in root, but got false.", dir1Name)
	}
	file1Name := "file1"
	if root.Exists(file1Name) {
		t.Errorf("'%s' doesn't exist in root, but got true.", file1Name)
	}
	entities.NewFile(file1Name, root)
	if !root.Exists(file1Name) {
		t.Errorf("'%s' exists in root, but got false.", file1Name)
	}

	// Test root's List()
	outputStr := fmt.Sprintf("[d] %s\n[f] %s\n", dir1Name, file1Name)
	if root.List() != outputStr {
		t.Errorf("Expected \n '%s' Got '%s'", outputStr, root.List())
	}

	// Test Remove() method.
	err := root.Remove(dir1Name)
	if err != nil {
		t.Errorf("Expected success removing '%s', but got error: '%s'.", dir1Name, err)
	}
	err = root.Remove(file1Name)
	if err != nil {
		t.Errorf("Expected success removing '%s', but got error: '%s'.", file1Name, err)
	}
	err = root.Remove(dir1Name)
	if err == nil {
		t.Error("Expected error, but got nothing.")
	}
	err = root.Remove(file1Name)
	if err == nil {
		t.Error("Expected error, but got nothing.")
	}

	// Test Rename() method.
	err = root.Rename("nie wolno")
	if err == nil {
		t.Errorf("Expected error renaming root, but got nothing.")
	}
}
