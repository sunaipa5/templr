package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func listTemplates() {
	files, err := os.ReadDir(templrDir())
	if err != nil {
		werror("Error reading templates: " + err.Error())
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
		warning("No template found")
	}
}

func initProject(templateName, projectName string) {
	src := filepath.Join(templrDir(), templateName)
	dest := filepath.Join(".", projectName)

	if _, err := os.Stat(src); os.IsNotExist(err) {
		werror("Template not found: " + templateName)
		return
	}

	err := copyDir(src, dest, "init.templr")
	if err != nil {
		werror("Error copying template: " + err.Error())
		return
	}

	winfo("Project '%s' created from template '%s'", projectName, templateName)

	if err := runInitTemplr(src, projectName); err != nil {
		werror("Error running init.templr: " + err.Error())
	}

	wsuccess("Done.")
}

func wsuccess(msg string, args ...any) {
	fmt.Printf("\033[32m%s\033[0m\n", fmt.Sprintf(msg, args...))
}

func werror(msg string, args ...any) {
	fmt.Printf("\033[31m%s\033[0m\n", fmt.Sprintf(msg, args...))
}

func winfo(msg string, args ...any) {
	fmt.Printf("\033[34m%s\033[0m\n", fmt.Sprintf(msg, args...))
}

func warning(msg string, args ...any) {
	fmt.Printf("\033[33m%s\033[0m\n", fmt.Sprintf(msg, args...))
}
