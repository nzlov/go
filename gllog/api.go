package gllog

import (
	log "github.com/Sirupsen/logrus"
	"github.com/yuin/gopher-lua"
)

var api = map[string]lua.LGFunction{
	"debugln": apiDebugln,
	"errorln": apiErrorln,
	"infoln":  apiInfoln,
}

func apiDebugln(L *lua.LState) int {
	n := L.GetTop()
	as := make([]interface{}, n)

	for i := 1; i <= n; i++ {
		any := L.CheckAny(i)
		as[i-1] = any
	}

	log.Debugln(as...)

	return 0
}
func apiErrorln(L *lua.LState) int {
	n := L.GetTop()
	as := make([]interface{}, n)

	for i := 1; i <= n; i++ {
		any := L.CheckAny(i)
		as[i-1] = any
	}

	log.Errorln(as...)
	return 0
}
func apiInfoln(L *lua.LState) int {
	n := L.GetTop()
	as := make([]interface{}, n)

	for i := 1; i <= n; i++ {
		any := L.CheckAny(i)
		as[i-1] = any
	}
	log.Infoln(as...)
	return 0
}
