package filemonitor

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type FileEntry struct {
	path         string
	name         string
	lastModified int64
	length       int64
	exists       bool
	directory    bool
}

func NewFileEntry(path string) (*FileEntry, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, errors.New("File Abs Path is missing!" + path)
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, errors.New("File Path is missing!" + path)
	}
	names := strings.Split(path, "\\")
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

func (this *FileEntry) Refresh() bool {
	fileInfo, err := os.Stat(this.path)
	if err != nil {
		return true
	}

	oldLastModified := this.lastModified
	oldLength := this.length
	oldDirectory := this.directory

	this.lastModified = fileInfo.ModTime().UnixNano()
	this.length = fileInfo.Size()
	this.directory = fileInfo.IsDir()

	return oldLastModified != this.lastModified || oldLength != this.length || oldDirectory != this.directory
}

func (this FileEntry) Path() string {
	return this.path
}
func (this FileEntry) Name() string {
	return this.name
}
func (this FileEntry) Directory() bool {
	return this.directory
}
func (this FileEntry) Level() int {
	return len(strings.Split(this.path, "\\")) - 1
}
func (this FileEntry) Length() int64 {
	return this.length
}
func (this FileEntry) LastModified() int64 {
	return this.lastModified
}
