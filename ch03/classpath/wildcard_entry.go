package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] // remove *
	var compositeEntry []Entry
	walkfFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}

		if strings.HasSuffix(path, "jar") || strings.HasSuffix(path, "JAR") {
			entry := newZipEntry(path)
			compositeEntry = append(compositeEntry, entry)
		}
		return nil
	}
	_ = filepath.Walk(baseDir, walkfFn)
	return compositeEntry
}
