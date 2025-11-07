package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func listTemplates() {
	files, err := os.ReadDir(templrDir())
	if err != nil {
		fmt.Println("Error reading templates:", err)
		return
	}

	if len(files) > 0 {
		fmt.Println("Available templates:")
		for _, f := range files {
			if f.IsDir() {
				fmt.Println(" -", f.Name())
			}
		}
	} else {
		fmt.Println("No template found")
	}
}

func initProject(templateName, projectName string) {
	src := filepath.Join(templrDir(), templateName)
	dest := filepath.Join(".", projectName)

	if _, err := os.Stat(src); os.IsNotExist(err) {
		fmt.Println("Template not found:", templateName)
		return
	}

	err := copyDir(src, dest)
	if err != nil {
		fmt.Println("Error copying template:", err)
		return
	}

	fmt.Printf("Project '%s' created from template '%s'\n", projectName, templateName)
}
