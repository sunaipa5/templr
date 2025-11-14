package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func runInitTemplr(templateDir, projectName string) error {
	initFile := filepath.Join(templateDir, "init.templr")
	projectDir := filepath.Join(".", projectName)

	if _, err := os.Stat(initFile); os.IsNotExist(err) {
		winfo("No init.templr found, skipping...")
		return nil
	}

	file, err := os.Open(initFile)
	if err != nil {
		return fmt.Errorf("failed to open init.templr: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		cmdStr := strings.ReplaceAll(line, "%p", projectName)

		parts := strings.Fields(cmdStr)
		if len(parts) == 0 {
			continue
		}

		cmd := exec.Command(parts[0], parts[1:]...)
		cmd.Dir = projectDir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if runtime.GOOS == "windows" {
			noCmd(cmd)
		}

		fmt.Printf("Running: %s\n", cmdStr)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("command failed: %s", cmdStr)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading init.templr: %w", err)
	}

	return nil
}
