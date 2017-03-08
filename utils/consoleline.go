package utils

import (
	"fmt"
	"strings"
	"sync"
)

var (
	consolelinenum = 0
	consolelinemux = sync.Mutex{}
)

func Printf(format string, o ...interface{}) {
	consolelinemux.Lock()
	fmt.Print(strings.Repeat("\b", consolelinenum))
	s := fmt.Sprintf(format, o...)
	consolelinenum = len(s)
	fmt.Print(s)
	consolelinemux.Unlock()
}
