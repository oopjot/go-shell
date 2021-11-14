package entities_test

import (
	"go-shell/entities"
	"testing"
)

func TestFile(t *testing.T) {
	root := entities.RootDir()

	testFileName := "test_file"
	file1, _ := entities.NewFile(testFileName, root)

	// Test file already exists
	_, err := entities.NewFile(testFileName, root)
	if err == nil {
		t.Error("Expected error, but got nothing.")
	}

	if file1.Name() != testFileName {
		t.Error("Expected true, but got false.")
	}

	if file1.Read() != "" {
		t.Errorf("Expected empty string, but got '%s'.", file1.Read())
	}

	writeText := "jeden"
	appendText := "dwa"
	expectedOne := "jeden\n"
	expectedTwo := "jeden\ndwa\n"
	file1.Write(writeText)
	if file1.Read() != expectedOne {
		t.Errorf("Expected '%s' Got '%s'", expectedOne, file1.Read())
	}
	file1.Append(appendText)
	if file1.Read() != expectedTwo {
		t.Errorf("Expected '%s' Got '%s'", expectedTwo, file1.Read())
	}
	file1.Write(writeText)
	if file1.Read() != expectedOne {
		t.Errorf("Expected '%s' Got '%s'", expectedOne, file1.Read())
	}

	err = file1.Rename("test_file_renamed")
	if err != nil {
		t.Errorf("Expected success, but got error: '%s'", err)
	}

}
