package warp

import (
	"os/exec"
)

func ExecWarp(logFunction func(msg string, args ...interface{}), cmd *exec.Cmd) *exec.Cmd {
	logFunction("exec: %s\n", cmd)
	return cmd
}
