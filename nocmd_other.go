//go:build !windows

package main

import "os/exec"

func noCmd(cmd *exec.Cmd) {

}
