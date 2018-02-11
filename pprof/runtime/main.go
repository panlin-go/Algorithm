package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func main() {
	pf, err := os.Create("test.prof")
	if err != nil {
		fmt.Println(err)
		return
	}
	pprof.StartCPUProfile(pf)

	pprof.StopCPUProfile()
}
