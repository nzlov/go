package glpool

import (
	"sync"

	"github.com/yuin/gopher-lua"
)

var (
	m       sync.Mutex
	lstates []*lua.LState
	n       int
)

func init() {
	m = sync.Mutex{}
	lstates = make([]*lua.LState, 0)
}
func Get() *lua.LState {
	m.Lock()
	defer m.Unlock()
	n := len(lstates)
	if n == 0 {
		return new()
	}
	x := lstates[n-1]
	lstates = lstates[0 : n-1]
	return x
}

func new() *lua.LState {
	L := lua.NewState()
	n++
	return L
}

func Size() int {
	return n
}

func Put(L *lua.LState) {
	m.Lock()
	defer m.Unlock()
	lstates = append(lstates, L)
}

func Shutdown() {
	for _, L := range lstates {
		L.Close()
	}
}
