package main

import (
	"fmt"
	"strings"
	"time"
	"unsafe"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	spew.Config.DisableMethods = true
	spew.Config.MaxDepth = 1
	spew.Config.Indent = ""

	tick := time.NewTicker(time.Second)
	for {
		t := <-tick.C
		print_time(t)
		ptr := unsafe.Pointer(&t)
		print_parts(
			*(*uint64)(ptr), // wall
			*(*int64)(unsafe.Pointer(uintptr(ptr) + uintptr(8))), // ext
			t.Location(),
		)
	}
}

//go:noinline
func print_time(time time.Time) {
	fmt.Println(strings.ReplaceAll(spew.Sdump(time), "\n", " "))
}

//go:noinline
func print_parts(wall uint64, ext int64, loc *time.Location) {
	fmt.Printf("(raw parts) { wall: (uint64) %d, ext: (int64) %d, loc: (*time.Location)(%p) }\n", wall, ext, loc)
}
