package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func showNumber(num int) {
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	fmt.Println(num, timestamp)
	time.Sleep(time.Millisecond * 1)
}

func main() {

	runtime.GOMAXPROCS(0)
	iterations := 10

	for i := 0; i <= iterations; i++ {
		go showNumber(i)
	}

	fmt.Println("GoodBye!")
	runtime.Gosched()
}
