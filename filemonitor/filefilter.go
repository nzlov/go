package filemonitor

import (
	"github.com/nzlov/go/utils"
	"strings"
)

type FileFilter struct {
	filterList *utils.Array
}

func NewFileFilter() *FileFilter {
	return &FileFilter{utils.NewArray()}
}

func (this *FileFilter) AddFilter(filter string) {
	this.filterList.Add(filter)
}

func (this *FileFilter) DelFilter(filter string) bool {
	return this.filterList.RemoveValue(filter)
}

func (this *FileFilter) SetFilters(filters ...string) {
	for _, filter := range filters {
		if filter != "" {
			this.AddFilter(filter)
		}
	}
}

func (this FileFilter) Check(path string) bool {
	for i := 0; i < this.filterList.Size(); i++ {
		v, err := this.filterList.Get(i)
		if err != nil {
			continue
		}
		switch inst := v.(type) {

		case string:

			if strings.HasSuffix(path, inst) {
				return true
			}
		}
	}
	return false
}
