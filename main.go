package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	showHidden := flag.Bool("a", false, "show files beginning with '.'")
	flag.Parse()

	lister := FileLister{
		showHidden: *showHidden,
	}
	console := ConsoleOutput{
		separator: "  ",
	}

	LS(lister, console)
}

func LS(lister Lister, output Output) {
	entries, err := lister.getEntries()
	if err != nil {
		output.addErr(err)
		return
	}

	for _, entry := range entries {
		output.addEntry(entry)
	}
	output.end()
}

type Lister interface {
	getEntries() ([]fs.DirEntry, error)
}

type Output interface {
	addErr(err error)
	addEntry(entry fs.DirEntry)
	end()
}

type FileLister struct {
	showHidden bool
}

func (f FileLister) getEntries() ([]fs.DirEntry, error) {
	entries, err := getAllEntries()
	if err != nil {
		return entries, err
	}

	result := []fs.DirEntry{}
	for _, entry := range entries {
		if f.showHidden || shouldShow(entry) {
			result = append(result, entry)
		}
	}

	return result, nil
}

func getAllEntries() ([]fs.DirEntry, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, errors.New("error finding working directory")
	}
	entries, err := os.ReadDir(pwd)
	if err != nil {
		return nil, fmt.Errorf("error listing directory: %s", pwd)
	}
	return entries, nil
}

func shouldShow(entry fs.DirEntry) bool {
	return entry.Name()[0] != '.'
}

type ConsoleOutput struct {
	separator string
}

func (c ConsoleOutput) addErr(err error) {
	fmt.Println(err.Error())
}

func (c ConsoleOutput) addEntry(entry fs.DirEntry) {
	print := getPrinter(entry)
	print(entry.Name(), c.separator)
}

func (c ConsoleOutput) end() {
	if c.separator != "\n" {
		fmt.Println()
	}
}

func getPrinter(entry fs.DirEntry) func(string, string) {
	if entry.IsDir() {
		return printDir
	}
	return printFile
}

func printFile(name string, separator string) {
	fmt.Printf("%s%s", name, separator)
}

func printDir(name string, separator string) {
	fmt.Printf("%s%s", yellow(name), separator)
}

func yellow(input string) string {
	return fmt.Sprintf("\u001b[33m%s\u001b[0m", input)
}
