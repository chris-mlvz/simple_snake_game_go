package main

import (
	"os"
	"os/exec"
	"runtime"
)

func clearWindows() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func clearLinux() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func clear() {
	if runtime.GOOS == "windows" {
		clearWindows()
	} else if runtime.GOOS == "linux" {
		clearLinux()
	}
}
