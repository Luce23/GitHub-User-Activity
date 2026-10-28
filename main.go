package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		printUsage()
		return
	}

	fmt.Println("GitHub User Activity CLI")
}

func printUsage() {
	fmt.Println("GitHub User Activity CLI")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  github-activity <username>")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -h, --help    Show help")
}
