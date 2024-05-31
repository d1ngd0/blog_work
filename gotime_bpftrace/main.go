package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	spew.Config.DisableMethods = true
	spew.Config.MaxDepth = 1
	spew.Config.Indent = ""

	tick := time.NewTicker(time.Second)
	for {
		<-tick.C
		print_time(time.Now())
	}
}

//go:noinline
func print_time(time time.Time) {
	fmt.Println(strings.ReplaceAll(spew.Sdump(time), "\n", " "))
}
