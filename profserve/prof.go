package profserve

import (
	"net/http"
	"net/http/pprof"
)

func StartProfServe(host string) {
	profServeMux := http.NewServeMux()
	profServeMux.HandleFunc("/debug/pprof/", pprof.Index)
	profServeMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	profServeMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	profServeMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	err := http.ListenAndServe(host, profServeMux)
	if err != nil {
		panic(err)
	}
}
