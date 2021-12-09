package main

import (
	"runtime"

	"github.com/debridge-finance/orbitdb-go/cli"
)

func init() { runtime.GOMAXPROCS(runtime.NumCPU()) }

func main() { cli.Run() }
