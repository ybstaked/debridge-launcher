package main

import (
	"fmt"
	"runtime"

	// XXX: uncomment to enable pprof
	"net/http"
	"net/http/pprof"

	"github.com/debridge-finance/orbitdb-go/app/orbitdb/cli"
)

func healthcheckHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func registerPProfHandlers(r *http.ServeMux) {
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)
}

func init() { runtime.GOMAXPROCS(runtime.NumCPU()) }

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/healthz", healthcheckHandler)
	registerPProfHandlers(m)

	// XXX: uncomment to enable pprof
	// go tool pprof -http=":8000" http://localhost:6060/debug/pprof/heap
	// go tool trace -http=':8000' http://localhost:6060/debug/pprof/trace
	go func() {
		panic(http.ListenAndServe("localhost:6060", m))
	}()
	cli.Run()

}
