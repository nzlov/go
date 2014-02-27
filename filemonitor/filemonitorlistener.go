package filemonitor

type FileMonitorListener interface {
	OnStart()
	FileCreate(file *FileEntry)
	FileModify(file *FileEntry)
	FileDelete(file *FileEntry)
	DirectoryCreate(file *FileEntry)
	DirectoryModify(file *FileEntry)
	DirectoryDelete(file *FileEntry)
	OnEnd()
}

type FileMonitorAdaptor struct {
}

func (this *FileMonitorAdaptor) OnStart() {

}
func (this *FileMonitorAdaptor) FileCreate(file *FileEntry) {

}
func (this *FileMonitorAdaptor) FileModify(file *FileEntry) {

}
func (this *FileMonitorAdaptor) FileDelete(file *FileEntry) {

}
func (this *FileMonitorAdaptor) DirectoryCreate(file *FileEntry) {

}
func (this *FileMonitorAdaptor) DirectoryModify(file *FileEntry) {

}
func (this *FileMonitorAdaptor) DirectoryDelete(file *FileEntry) {

}
func (this *FileMonitorAdaptor) OnEnd() {

}
