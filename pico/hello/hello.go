package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"time"
)

func scan_char(ch chan string) {
	buf := make([]byte, 1)
	for {
		if _, err := io.ReadFull(os.Stdin, buf); err != nil {
			fmt.Printf("io.ReadAll Error: %s\n", err)
		}
		fmt.Println("double")
		ch <- string(buf)
	}
}

func main() {

	c1 := make(chan string, 1)
	go scan_char(c1)

	var i uint64 = 0
	for {
		select {
		case _ = <-c1:
			i = i * 2
			if i > math.MaxUint64/2 {
				i = 0
			}
		case <-time.After(100 * time.Millisecond):
			fmt.Printf("Hello TinyGo!: %d\n", i)
			i++
		}
	}
}
