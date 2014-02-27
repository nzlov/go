package filemonitor

import (
	"time"
)

type FileMonitor struct {
	running  bool
	observer *FileMonitorObserver
	dt       time.Duration
}

func NewFileMonitor(observer *FileMonitorObserver) *FileMonitor {
	return &FileMonitor{
		running:  false,
		observer: observer,
		dt:       time.Second}
}
func NewFileMonitorByDt(observer *FileMonitorObserver, dt time.Duration) *FileMonitor {
	return &FileMonitor{
		running:  false,
		observer: observer,
		dt:       dt}
}

func (this *FileMonitor) Start() {
	this.running = true
	go func() {
		for this.running {
			if this.observer != nil {
				this.observer.Check()
			}
			time.Sleep(this.dt)
		}
	}()
}
func (this *FileMonitor) End() {
	this.running = false
}
