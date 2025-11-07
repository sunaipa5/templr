package main

import (
	"fmt"
	"os"
)

func main() {
	os.MkdirAll(templrDir(), 0755)

	if len(os.Args) < 2 {
		fmt.Println("Usage: templr [list|init]")
		return
	}

	switch os.Args[1] {
	case "list":
		listTemplates()
	case "init":
		if len(os.Args) < 4 {
			fmt.Println("Usage: templr init <templateName> <projectName>")
			return
		}
		initProject(os.Args[2], os.Args[3])
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
