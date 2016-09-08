package glpool

import (
	"sync"

	"github.com/yuin/gopher-lua"
)

type Pool struct {
	m       sync.Mutex
	lstates []*lua.LState
	opt     lua.Options
}

func NewPool(opt lua.Options) *Pool {
	p := &Pool{
		m:       sync.Mutex{},
		lstates: make([]*lua.LState, 0),
		opt:     opt,
	}
	return p
}

func NewDefaultPool() *Pool {
	p := &Pool{
		m:       sync.Mutex{},
		lstates: make([]*lua.LState, 0),
		opt: lua.Options{
			CallStackSize: lua.CallStackSize,
			RegistrySize:  lua.RegistrySize,
		},
	}
	return p
}

func (self *Pool) Get() *lua.LState {
	self.m.Lock()

	n := len(self.lstates)
	if n == 0 {
		self.m.Unlock()
		return lua.NewState(self.opt)
	}
	x := self.lstates[n-1]
	self.lstates = self.lstates[0 : n-1]
	self.m.Unlock()
	return x
}

func (self *Pool) Put(l *lua.LState) {
	self.m.Lock()
	self.lstates = append(self.lstates, l)
	self.m.Unlock()
}

func (self *Pool) Size() int {
	self.m.Lock()
	i := len(self.lstates)
	self.m.Unlock()
	return i
}

func (self *Pool) Close() {
	self.m.Lock()
	for _, l := range self.lstates {
		l.Close()
	}
	self.lstates = make([]*lua.LState, 0)
	self.m.Unlock()
}
