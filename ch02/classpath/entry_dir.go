package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	abs, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{abs}
}

func (self *DirEntry)readClass(className string) ([]byte, Entry, error)  {
	filename := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(filename)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
