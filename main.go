package main

import (
	"fmt"
	"go-shell/entities"
)

func main() {

	fmt.Println("Siema")

	root := entities.RootDir()

	dir1, _ := entities.NewDir("dir1", root)
	fmt.Println(dir1.Name())

	file1, _ := entities.NewFile("file1", dir1)
	fmt.Println(file1.Name())

	fmt.Print(dir1.List())

	fmt.Print(file1.Read())
	file1.Write("siema elo yo")
	file1.Append("witam")
	file1.Append("pytam")
	fmt.Print(file1.Read())
	file1.Write("reset")
	fmt.Print(file1.Read())

	fmt.Println(dir1.Exists("file1"))
	dir1.Remove("file1")
	fmt.Println(dir1.Exists("file1"))
	entities.NewFile("file2", dir1)
	_, err := entities.NewFile("file2", dir1)
	if err != nil {
		fmt.Println(err)
	}

	entities.NewDir("child_to_dir1", dir1)
	fmt.Print(dir1.List())

}
