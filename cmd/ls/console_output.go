package ls

import (
	"fmt"
	"io/fs"
)

type ConsoleOutput struct {
	Separator string
}

func (c ConsoleOutput) AddErr(err error) {
	fmt.Println(err.Error())
}

func (c ConsoleOutput) AddEntry(entry fs.DirEntry) {
	print := getPrinter(entry)
	print(entry.Name(), c.Separator)
}

func (c ConsoleOutput) End() {
	if c.Separator != "\n" {
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
