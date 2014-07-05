package filemonitor

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func NewFileEntry(path string) (*FileEntry, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, errors.New("File Abs Path is missing!" + path)
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, errors.New("File Path is missing!" + path)
	}
	names := strings.Split(path, "/")
	name := names[len(names)-1]
	fileEntry := &FileEntry{
		path:         path,
		name:         name,
		lastModified: fileInfo.ModTime().UnixNano(),
		length:       fileInfo.Size(),
		exists:       true,
		directory:    fileInfo.IsDir()}
	return fileEntry, nil
}
