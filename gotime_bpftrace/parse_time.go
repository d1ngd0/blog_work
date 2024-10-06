package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
	"unsafe"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanLines)

	var buf struct {
		Wall     uint64 `json:"wall"`
		Ext      int64  `json:"ext"`
		location *time.Location
	}

	for s.Scan() {
		// unmarshal the data into the buff
		err := json.Unmarshal(s.Bytes(), &buf)
		if err != nil {
			// if we can't unmarshal then output it
			fmt.Println(s.Text())
			continue
		}

		// buf has the same in memory layout of time.Time, so we can
		// do an unsafe type conversion to turn it into a time.Time
		ti := *(*time.Time)(unsafe.Pointer(&buf))
		fmt.Println(ti.String())
	}
}
