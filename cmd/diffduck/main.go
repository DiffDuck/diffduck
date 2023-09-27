package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Runing diffduck")
	if len(os.Args) < 2 {
		fmt.Println("Usage: diffduck <path>")
		os.Exit(1)
	}

	arg := os.Args[1]
	if arg == "-v" || arg == "--version" {
		fmt.Println("DiffDuck v0.1.0")
		os.Exit(0)
	}

	path := filepath.Clean(arg)

	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	if fileInfo.IsDir() {
		fmt.Println("Error: ", path, "is a directory, not a file.")
		os.Exit(1)
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" && line[0:1] != "#" {
			fmt.Println("Commit message is not empty. Skipping.")
			os.Exit(0)
		}
	}

	fmt.Println("Writing commit message to", path)
	if err := os.WriteFile(path, []byte("Hello, DiffDuck!\n"), 0644); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	os.Exit(0)
}
