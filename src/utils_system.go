package main

import (
	"os/exec"
	"strings"
)

func Execute(command string) (bool, string, error) {
	 // Use a shell to execute the command
	 out, err := exec.Command("sh", "-c", command).Output()
	 if err != nil {
		 return false, string(out), err
	 }
	 return true, string(out), nil
 }
 
