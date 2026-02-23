package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "add":
		// handle add
		fmt.Println("add called")
	case "ls":
		// handle ls
		fmt.Println("ls called")
	case "find":
		// handle find
		fmt.Println("find called")
	case "rm":
		// handle rm
		fmt.Println("rm called")
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`cheat - a CLI command notebook

Usage:
    cheat add <command> [-t tag] [-d description]
    cheat ls [-t tag]
    cheat find <query>
    cheat rm <id>`)
}
