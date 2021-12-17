package main

import (
	"runtime"

	"github.com/debridge-finance/orbitdb-go/app/orbitdb/cli"
)

func init() { runtime.GOMAXPROCS(runtime.NumCPU()) }

func main() { cli.Run() }
