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

	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "Error: GitHub username is required.")
		fmt.Fprintln(os.Stderr, "Run 'github-activity --help' for usage.")
		os.Exit(1)
	}

	if len(os.Args) > 2 {
		fmt.Fprintln(os.Stderr, "Error: too many arguments.")
		fmt.Fprintln(os.Stderr, "Run 'github-activity --help' for usage.")
		os.Exit(1)
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
