package main

import (
	"os/exec"
	"strings"
)

type CommandResult struct {
    Success bool
    Output  string
    Error   error
}

func Execute(command string) CommandResult {
    cmd := exec.Command("sh", "-c", command)
    output, err := cmd.CombinedOutput()

    result := CommandResult{
        Success: err == nil,
        Output:  string(output),
        Error:   err,
    }

    return result
}
