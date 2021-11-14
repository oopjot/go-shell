package commands

import (
	"os"
	"os/exec"
	"runtime"
	"errors"
)

func Clear() error {
	var cmd *exec.Cmd
	if runtime.GOOS == "linux" {
		cmd = exec.Command("clear")
	} else if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		return errors.New("clear: OS not supported, can't clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil
}
