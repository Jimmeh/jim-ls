package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	all := flag.Bool("a", false, "show all files/folders")
	flag.Parse()

	fmt.Printf("displaying all files: %v", *all)
	fmt.Println()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Unable to get current directory")
		return
	}
	fmt.Printf("displaying contents of %v", dir)
	fmt.Println()

	listing, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Unable to get contents of current directory")
		return
	}

	for _, entry := range listing {
		entryLine := entry.Name()
		if *all || entryLine[0] != '.' {
			fmt.Println(entryLine)
		}
	}
}
