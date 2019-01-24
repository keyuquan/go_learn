package main

import (
	"fmt"
	"runtime"
	"time"
)

func test() {
	defer fmt.Println("ccccccccccccccccc")
	//return
	runtime.Goexit() // 退出当前go程。
	defer fmt.Println("ddddddddddddddddd")
}

func main() {

	go func() {
		defer fmt.Println("aaaaaaaaaa")
		test()
		fmt.Println("bbbbbbbbbbbbb")
	}()

	for {
		time.Sleep(100 * time.Millisecond)
	}
}
