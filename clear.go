package main

import (
	"os"
	"os/exec"
)

func clearScreen(){
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}