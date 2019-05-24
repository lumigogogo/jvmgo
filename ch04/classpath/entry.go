package classpath

import (
	"os"
	"strings"
)

const parseListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(classname string) (bytes []byte, entry Entry, err error)

	String() string
}

func newEntry(path string) Entry {

	if strings.Contains(path, parseListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, "jar") || (strings.HasSuffix(path, "JAR")) ||
		strings.HasSuffix(path, "zip") || strings.HasSuffix(path, "ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
