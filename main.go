package main

import (
	"github.com/taion809/gogarithms/commands"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	commands.Execute()
}
