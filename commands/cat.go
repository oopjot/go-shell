package commands

import (
	"go-shell/entities"
	"go-shell/utils"
	"bufio"
	"fmt"
	"os"
)

func Cat(currDir entities.Dir, input ...string) error {
	if len(input) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fmt.Println(scanner.Text())
		return nil
	}

	var content string
	for _, path := range input {
		var file entities.File
		dest, filename := utils.GetDest(path)
		dir, err := utils.Unpath(dest, currDir)
		if err != nil {
			return err
		}
		file, err = dir.FindFile(filename)
		if err != nil {
			return err
		}
		content = content + file.Read()
	}
	fmt.Print(content)
	return nil
}
