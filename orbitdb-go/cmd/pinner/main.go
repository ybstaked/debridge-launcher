package main

import (
	"runtime"

	"github.com/debridge-finance/orbitdb-go/app/pinner/cli"
)

func init() { runtime.GOMAXPROCS(runtime.NumCPU()) }

func main() { cli.Run() }
