package main

import (
    "os/exec"
    "runtime"
)

type CommandResult struct {
    Success bool
    Output  string
    Error   error
}

func Execute(command string) CommandResult {
    var cmd *exec.Cmd
    if runtime.GOOS == "windows" {
        cmd = exec.Command("cmd", "/C", command)
    } else {
        cmd = exec.Command("sh", "-c", command)
    }

    output, err := cmd.CombinedOutput()

    result := CommandResult{
        Success: err == nil,
        Output:  string(output),
        Error:   err,
    }

    return result
}
