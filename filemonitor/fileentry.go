package filemonitor

import (
	"os"
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
