package classpath

import (
	"errors"
	"strings"
)

// 复合
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	var compositeEntry []Entry
	for _, path := range strings.Split(pathList, parseListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (self CompositeEntry) readClass(classname string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.readClass(classname)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found " + classname)
}

func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, parseListSeparator)
}
