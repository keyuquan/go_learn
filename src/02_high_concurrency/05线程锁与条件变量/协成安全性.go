package main

import (
	"fmt"
	"runtime"
	"time"
)

var a = 1

func test6() {
	for i := 0; i < 100000; i++ {
		a = a + 1
		fmt.Println("test6==============", a)
	}
}

func test5() {
	for i := 0; i < 100000; i++ {
		fmt.Println("test5+++", a)
		a = a + 1
		fmt.Println("test5++++++++++++++", a)
	}
}

func main() {

	go test5()
	go test6()
	for {
		runtime.Gosched() // 出让当前 cpu 时间片。
		time.Sleep(1000)
	}

}
