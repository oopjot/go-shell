package entities_test

import (
	"go-shell/entities"
	"testing"
	"fmt"
)

func TestDir(t *testing.T) {
	root := entities.RootDir()
	testDir, _ := entities.NewDir("test_dir", root)

	// Assert that testDir dir name is "/"
	if testDir.Name() != "test_dir" {
		t.Errorf("testDir's name incorrect. Should be '/', got '%s'.", testDir.Name())
	}

	// Test isRoot()
	if testDir.IsRoot() {
		t.Error("testDir is not root, but got true")
	}

	// Test testDir's Exists() method.
	dir1Name := "dir1"
	if testDir.Exists(dir1Name) {
		t.Errorf("'%s' doesn't exist in testDir, but got true.", dir1Name)
	}
	entities.NewDir(dir1Name, testDir)
	if !testDir.Exists(dir1Name) {
		t.Errorf("'%s' exists in testDir, but got false.", dir1Name)
	}
	file1Name := "file1"
	if testDir.Exists(file1Name) {
		t.Errorf("'%s' doesn't exist in testDir, but got true.", file1Name)
	}
	entities.NewFile(file1Name, testDir)
	if !testDir.Exists(file1Name) {
		t.Errorf("'%s' exists in testDir, but got false.", file1Name)
	}

	// Test testDir's List()
	outputStr := fmt.Sprintf("[d] %s\n[f] %s\n", dir1Name, file1Name)
	if testDir.List() != outputStr {
		t.Errorf("Expected \n '%s' Got '%s'", outputStr, testDir.List())
	}

	// Test Remove() method.
	err := testDir.Remove(dir1Name)
	if err != nil {
		t.Errorf("Expected success removing '%s', but got error: '%s'.", dir1Name, err)
	}
	err = testDir.Remove(file1Name)
	if err != nil {
		t.Errorf("Expected success removing '%s', but got error: '%s'.", file1Name, err)
	}
	err = testDir.Remove(dir1Name)
	if err == nil {
		t.Error("Expected error, but got nothing.")
	}
	err = testDir.Remove(file1Name)
	if err == nil {
		t.Error("Expected error, but got nothing.")
	}

	// Test Rename() method.
	newName := "test_dir_renamed"
	err = testDir.Rename(newName)
	if err != nil {
		t.Errorf(
			"Expected success renaming '%s' to '%s', nunt got error: '%s'",
			testDir.Name(),
			newName,
			err,
		)
	}
	entities.NewDir("dir1", root)
	err = testDir.Rename("dir1")
	if err == nil {
		t.Errorf("Expected error renaming testDir, but got nothing.")
	}
}
