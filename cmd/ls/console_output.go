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
	if entry.IsDir() {
		printDir(entry.Name(), c.Separator)
	}
	printFile(entry.Name(), c.Separator)
}

func (c ConsoleOutput) End() {
	if c.Separator != "\n" {
		fmt.Println()
	}
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
