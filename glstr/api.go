package glstr

import (
	"github.com/layeh/gopher-luar"
	"github.com/yuin/gopher-lua"
	"strings"
)

var api = map[string]lua.LGFunction{
	"split": apiSplit,
}

func apiSplit(L *lua.LState) int {
	str := L.CheckString(1)
	d := L.CheckString(2)

	ns := strings.Split(str, d)
	L.Push(luar.New(L, ns))
	return 1
}
