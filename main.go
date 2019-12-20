package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	bi, ok := debug.ReadBuildInfo()
	fmt.Printf("%v\n%+v\n", ok, bi)
}
