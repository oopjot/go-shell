package utils_test

import (
	"go-shell/entities"
	"go-shell/utils"
	"testing"
)

func TestUtils (t *testing.T) {
	root := entities.RootDir()
	dir1, _ := entities.NewDir("dir1", root)
	dir2, _ := entities.NewDir("dir2", dir1)
	dir3, _ := entities.NewDir("dir3", dir2)
	found, err := utils.GetRoot(dir3)
	if err != nil {
		t.Errorf("Expected true, got error: '%s'", err)
	}
	if !found.IsRoot() {
		t.Errorf("Expected true, got false. %s", found.Name())
	}

	// Test Unpath with abs path
	found, err = utils.Unpath("/dir1/dir2/dir3", dir2)
	if err != nil {
		t.Errorf("Unexpected error occured: '%s'", err)
	}
	if found.Name() != "dir3" {
		t.Errorf("'dir3' expected, got '%s'", found.Name())
	}

	// Test Unpath with some relative path
	found, err = utils.Unpath("dir2/dir3", dir1)
	if err != nil {
		t.Errorf("Unexpected error occured: '%s'", err)
	}
	if found.Name() != "dir3" {
		t.Errorf("'dir3' expected, got '%s'", found.Name())
	}

	found, err = utils.Unpath("./dir2/dir3", dir1)
	if err != nil {
		t.Errorf("Unexpected error occured: '%s'", err)
	}
	if found.Name() != "dir3" {
		t.Errorf("'dir3' expected, got '%s'", found.Name())
	}

	found, err = utils.Unpath("../dir1/dir2/dir3", dir1)
	if err != nil {
		t.Errorf("Unexpected error occured: '%s'", err)
	}
	if found.Name() != "dir3" {
		t.Errorf("'dir3' expected, got '%s'", found.Name())
	}

	// Test Unpath for errors
	found, err = utils.Unpath("/dir2/dir3", dir1)
	if err == nil {
		t.Errorf("Error expected.")
	}

	found, err = utils.Unpath("../dir2/dir3", dir1)
	if err == nil {
		t.Errorf("Error expected.")
	}

	found, err = utils.Unpath("../././dir3", dir1)
	if err == nil {
		t.Errorf("Error expected.")
	}

}
