package utils

import (
	"container/list"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"runtime"
	"strings"
)

const (
	Number    = "0123456789"
	LowerCase = "abcdefghijklmnopqrstuvwxyz"
	UpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Symbol    = "~!@#$%^&*()_+-=/."
)

type MD5CrackOption struct {
	GoNum        int
	PasswdLenMin int
	PasswdLenMax int
	CacheChNum   int
}

type MD5CrackReturnType int

const (
	MD5CrackReturnType_Finish = iota
	MD5CrackReturnType_Error
	MD5CrackReturnType_Check
)

type MD5CrackReturn struct {
	GoID   int
	Status MD5CrackReturnType
	Recive string
}

func NewMD5CrackOption() *MD5CrackOption {
	return &MD5CrackOption{
		GoNum:        runtime.NumCPU() * 2,
		PasswdLenMin: 1,
		PasswdLenMax: 0,
		CacheChNum:   50,
	}
}

type MD5Crack struct {
	keyString string
	isRunning bool
	option    *MD5CrackOption
	keyIndex  *list.List
}

func NewMD5Crack(o *MD5CrackOption) *MD5Crack {
	md5crack := &MD5Crack{
		option: o,
	}
	return md5crack
}

func (m *MD5Crack) Start(passwd, k string) chan MD5CrackReturn {
	recive := make(chan MD5CrackReturn, m.option.CacheChNum*2)
	keymsg := make(chan string, m.option.CacheChNum)

	exit := make(chan bool)

	m.keyIndex = list.New()
	m.keyString = k
	m.isRunning = true

	go m.start(keymsg, recive, exit)

	for i := 0; i <= m.option.GoNum; i++ {

		go m.check(i, passwd, keymsg, recive, exit)
	}

	return recive
}

func (m *MD5Crack) generatekey() (string, error) {
	if m.option.PasswdLenMax > 0 && m.keyIndex.Len() > m.option.PasswdLenMax {
		return "", errors.New("Beyond the maximum number of digits!")
	}
	s := ""
	v := m.keyIndex.Front()
	for v != nil {
		s = string(m.keyString[v.Value.(int)]) + s
		v = v.Next()
	}
	m.add(0)
	return string(s), nil
}

func (m *MD5Crack) add(index int) {
	v := m.keyIndex.Front()
	for i := 0; i < index; i++ {
		v = v.Next()
	}
	if v == nil {
		m.keyIndex.PushBack(0)
	} else {
		vi := v.Value.(int)
		if vi+1 == len(m.keyString) {
			v.Value = 0
			m.add(index + 1)
		} else {
			v.Value = vi + 1
		}
	}
}

func (m *MD5Crack) start(keyMsg chan string, over chan MD5CrackReturn, exit chan bool) {
	// defer fmt.Println("MD5Crack start end")
	for m.isRunning {
		s := ""
		for i := 0; i < m.option.CacheChNum/2; i++ {
			k, err := m.generatekey()
			if err != nil {
				if i == 0 {
					m.isRunning = false
					close(keyMsg)
					over <- MD5CrackReturn{
						GoID:   -1,
						Status: MD5CrackReturnType_Error,
						Recive: err.Error(),
					}
					return
				} else {
					break
				}
			}
			s = s + ";" + k
		}
		select {
		case keyMsg <- s:
		case <-exit:
			close(keyMsg)
			return
		}
	}
}

func (m *MD5Crack) check(id int, mw string, keyMsg chan string, over chan MD5CrackReturn, exit chan bool) {
	h := md5.New()
	// defer fmt.Printf("MD5Crack check %d end\n", id)

	for keys := range keyMsg {
		key := strings.Split(keys, ";")
		for _, value := range key {
			if !m.isRunning {
				return
			}
			h.Reset()
			h.Write([]byte(value))
			nm := hex.EncodeToString(h.Sum(nil))
			over <- MD5CrackReturn{
				GoID:   id,
				Status: MD5CrackReturnType_Check,
				Recive: value,
			}
			if nm == mw {
				over <- MD5CrackReturn{
					GoID:   id,
					Status: MD5CrackReturnType_Finish,
					Recive: value,
				}
				m.isRunning = false
				exit <- true
				return
			}
		}
	}
}
