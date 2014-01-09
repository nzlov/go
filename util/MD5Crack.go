package main

import (
	"container/list"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"time"
)

type Key struct {
	minl      int
	maxl      int
	keyString string
	keyIndex  *list.List
}

//k 查询字符表
func NewKey(k string) *Key {
	return NewMinKey(k, 1)
}

//k 查询字符表 minl最短位数
func NewMinKey(k string, minl int) *Key {
	return NewMinMaxKey(k, minl, 0)
}

//k 查询字符表 maxl最长位数
func NewMaxKey(k string, maxl int) *Key {
	return NewMinMaxKey(k, 1, maxl)
}

//k 查询字符表 minl最长位数 maxl最长位数
func NewMinMaxKey(k string, minl, maxl int) *Key {
	key := &Key{}
	key.minl = minl
	key.maxl = maxl
	key.keyString = k
	key.Init()
	return key
}

//初始化
func (k *Key) Init() {
	k.keyIndex = list.New()
	for i := 0; i < k.minl; i++ {
		k.keyIndex.PushBack(0)
	}
}

//生成匹配key
func (k *Key) Generate() (string, error) {
	if k.maxl > 0 && k.keyIndex.Len() > k.maxl {
		return "", errors.New("Beyond the maximum number of digits!")
	}
	s := ""
	v := k.keyIndex.Front()
	for v != nil {
		s = string(k.keyString[v.Value.(int)]) + s
		v = v.Next()
	}
	k.add(0)
	return string(s), nil
}

func (k *Key) add(index int) {
	v := k.keyIndex.Front()
	for i := 0; i < index; i++ {
		v = v.Next()
	}
	if v == nil {
		k.keyIndex.PushBack(0)
	} else {
		vi := v.Value.(int)
		if vi+1 == len(k.keyString) {
			v.Value = 0
			k.add(index + 1)
		} else {
			v.Value = vi + 1
		}
	}
}

//匹配字符串生成
func generate(keyString string, min, max, ckNum int, keyMsg chan string, msg, errOver chan bool) {
	key := NewMinMaxKey(keyString, min, max)
	for {
		<-msg
		s := ""
		for i := 0; i < ckNum; i++ {
			k, err := key.Generate()
			if err != nil {
				if i == 0 {
					errOver <- true
					return
				} else {
					break
				}
			}
			s = s + ";" + k
		}
		keyMsg <- s
	}
}

//检测
func check(mw string, keyMsg chan string, msg chan bool, over chan string) {
	for {
		msg <- true
		keys := <-keyMsg
		key := strings.Split(keys, ";")
		for _, value := range key {
			h := md5.New()
			h.Write([]byte(value)) // 需要加密的字符串为 123456
			m := hex.EncodeToString(h.Sum(nil))
			fmt.Println("check:", value)
			if m == mw {
				over <- value
				return
			}
		}
	}
}

func main() {
	t1 := time.Now()
	runtime.GOMAXPROCS(runtime.NumCPU())
	keyString := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()_+-=/." //检测字符表
	mw := "b28b7d691bb595df727a47cbf1240464"                                                       //需要破解的密文 明文：008784
	num := 10                                                                                      //开启线程数
	min := 1                                                                                       //最短位数
	max := 20                                                                                      //最长位数
	ckNum := 20                                                                                    //每次生成匹配数
	keyMsg := make(chan string)                                                                    //传递生成检测key
	msg := make(chan bool)                                                                         //传递需要生成检测key
	over := make(chan string)                                                                      //完成
	errOver := make(chan bool)                                                                     //未完成

	go generate(keyString, min, max, ckNum, keyMsg, msg, errOver)

	for i := 0; i < num; i++ {
		go check(mw, keyMsg, msg, over)
	}

	select {
	case v := <-over:
		fmt.Println("find:", v, "time:", time.Now().Sub(t1))
	case <-errOver:
		fmt.Println("Not find! time:", time.Now().Sub(t1))
	}
}
