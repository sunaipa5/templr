//go:build windows

package main

import (
	"os/exec"
	"syscall"
)

func noCmd(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
}
