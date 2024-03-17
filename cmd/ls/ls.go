package ls

import "io/fs"

func Ls(lister Lister, output Output) {
	entries, err := lister.GetEntries()
	if err != nil {
		output.AddErr(err)
		return
	}

	for _, entry := range entries {
		output.AddEntry(entry)
	}
	output.End()
}

type Lister interface {
	GetEntries() ([]fs.DirEntry, error)
}

type Output interface {
	AddErr(err error)
	AddEntry(entry fs.DirEntry)
	End()
}
