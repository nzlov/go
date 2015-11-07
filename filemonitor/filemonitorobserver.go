package filemonitor

import (
	"github.com/nzlov/go/array"
	"os"
	"path/filepath"
)

type FileMonitorObserver struct {
	fileListeners *array.Array
	fileMap       map[string]*FileEntry
	tempfileMap   *array.Array
	one           bool
	rootEntry     *FileEntry
	fileFileter   *FileFilter
}

func NewFileMonitorObserver(path string) *FileMonitorObserver {
	fileEntry, _ := NewFileEntry(path)
	return NewFileMonitorObserverByFileEntry(fileEntry)
}

func NewFileMonitorObserverByFileFilter(path string, fileter *FileFilter) *FileMonitorObserver {
	fileEntry, _ := NewFileEntry(path)
	return NewFileMonitorObserverByFileEntryAndFileFileter(fileEntry, fileter)
}

func NewFileMonitorObserverByFileEntry(rootEntry *FileEntry) *FileMonitorObserver {
	return NewFileMonitorObserverByFileEntryAndFileFileter(rootEntry, nil)
}
func NewFileMonitorObserverByFileEntryAndFileFileter(rootEntry *FileEntry, fileter *FileFilter) *FileMonitorObserver {
	observer := &FileMonitorObserver{
		fileListeners: array.NewArray(),
		fileMap:       make(map[string]*FileEntry),
		tempfileMap:   array.NewArray(),
		one:           true,
		rootEntry:     rootEntry,
		fileFileter:   fileter}
	return observer
}

func (this *FileMonitorObserver) AddListener(listener FileMonitorListener) {
	this.fileListeners.Add(listener)
}

func (this *FileMonitorObserver) DelListener(listener FileMonitorListener) {
	this.fileListeners.RemoveValue(listener)
}

func (this *FileMonitorObserver) Check() {
	for i := 0; i < this.fileListeners.Size(); i++ {
		v, err := this.fileListeners.Get(i)
		if err != nil {
			continue
		}
		switch inst := v.(type) {
		case FileMonitorListener:
			inst.OnStart()
		}
	}

	this.tempfileMap.Clear()

	//判断是否有删除的文件或目录
	for k, file := range this.fileMap {
		if _, err := os.Stat(k); err != nil {
			delete(this.fileMap, k)
			for i := 0; i < this.fileListeners.Size(); i++ {
				v, err := this.fileListeners.Get(i)
				if err != nil {
					continue
				}
				switch inst := v.(type) {
				case FileMonitorListener:
					if file.Directory() {
						inst.DirectoryDelete(file)
					} else {
						inst.FileDelete(file)
					}
				}
			}
		}
	}

	filepath.Walk(this.rootEntry.Path(), func(path string, fi os.FileInfo, err error) error {
		if nil == fi {
			return err
		}
		// if fi.IsDir() {
		// 	return nil
		// }
		// name := fi.Name()

		if this.fileFileter != nil {
			if !this.fileFileter.Check(path) && !fi.IsDir() {
				return nil
			}
		}

		//第一次监测时只添加
		if this.one {
			fileEntry, err := NewFileEntry(path)
			if err != nil {
				return nil
			}
			this.fileMap[path] = fileEntry
			return nil
		}

		this.tempfileMap.Add(path)

		//判断是否新建
		file, ok := this.fileMap[path]
		if !ok {
			fileEntry, err := NewFileEntry(path)
			if err != nil {
				return nil
			}
			this.fileMap[path] = fileEntry
			for i := 0; i < this.fileListeners.Size(); i++ {
				v, err := this.fileListeners.Get(i)
				if err != nil {
					continue
				}
				switch inst := v.(type) {
				case FileMonitorListener:
					if fi.IsDir() {
						inst.DirectoryCreate(fileEntry)
					} else {
						inst.FileCreate(fileEntry)
					}
				}
			}
			return nil
		}

		//如果不是新建就监测是否修改
		if file.Refresh() {
			for i := 0; i < this.fileListeners.Size(); i++ {
				v, err := this.fileListeners.Get(i)
				if err != nil {
					continue
				}
				switch inst := v.(type) {
				case FileMonitorListener:
					if fi.IsDir() {
						inst.DirectoryModify(file)
					} else {
						inst.FileModify(file)
					}
				}
			}
			return nil
		}
		return nil
	})

	if this.one {
		this.one = false
	} else {

	}

	for i := 0; i < this.fileListeners.Size(); i++ {
		v, err := this.fileListeners.Get(i)
		if err != nil {
			continue
		}
		switch inst := v.(type) {
		case FileMonitorListener:
			inst.OnEnd()
		}
	}
}
