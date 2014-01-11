// check.go
package check

import (
	"fmt"
	"net/http"
	"time"
)

type EKCheck struct {
	Url       string
	Interval  time.Duration
	EMail     EKMail
	isRunning bool
}

func (ekc *EKCheck) Start() {
	ekc.isRunning = true
	go ekc.check()
}

func (ekc *EKCheck) Stop() {
	ekc.isRunning = false
}

func (ekc *EKCheck) check() {
	isNew := false
	for ekc.isRunning {
		_, err := http.Get(ekc.Url)
		if err != nil && !isNew {
			mErr := ekc.EMail.SendMail()
			if mErr != nil {
				fmt.Println(mErr.Error())
			}
			isNew = true
		} else if err == nil && isNew {
			isNew = false
		}
		time.Sleep(ekc.Interval * time.Millisecond)
	}
}

func (ekc *EKCheck) setUrl(url string) {
	ekc.Url = url
}

func (ekc *EKCheck) setInterval(interval int) {
	ekc.Interval = time.Duration(interval)
}

func (ekc *EKCheck) setEMail(eMail EKMail) {
	ekc.EMail = eMail
}

func (ekc *EKCheck) getEMail() EKMail {
	return ekc.EMail
}
