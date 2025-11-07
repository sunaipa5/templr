package main

import (
	"os"
	"path/filepath"
	"runtime"
)

func templrDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	var documents string
	switch runtime.GOOS {
	case "windows":
		documents = filepath.Join(home, "Documents")
	default:
		documents = filepath.Join(home, "Documents")
	}

	return filepath.Join(documents, "templr")
}
