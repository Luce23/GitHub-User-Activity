package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
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

	username := strings.TrimSpace(os.Args[1])
	if err := validateUsername(username); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v.\n", err)
		os.Exit(1)
	}

	fmt.Println("GitHub User Activity CLI")
}

func validateUsername(username string) error {
	if username == "" {
		return errors.New("username cannot be empty")
	}
	if len(username) > 39 {
		return errors.New("username must not exceed 39 characters")
	}
	if strings.HasPrefix(username, "-") || strings.HasSuffix(username, "-") {
		return errors.New("username cannot start or end with a hyphen")
	}
	if strings.Contains(username, "--") {
		return errors.New("username cannot contain consecutive hyphens")
	}
	for _, char := range username {
		isLowercaseLetter := char >= 'a' && char <= 'z'
		isUppercaseLetter := char >= 'A' && char <= 'Z'
		isDigit := char >= '0' && char <= '9'

		if !isLowercaseLetter && !isUppercaseLetter && !isDigit && char != '-' {
			return errors.New("username can contain only letters, numbers, and hyphens")
		}
	}
	return nil
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
