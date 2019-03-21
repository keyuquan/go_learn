package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(out chan<- int, idx int) {
	for i := 0; i < 50; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("生产者%dth，生产：%d\n", idx, i)
		out <- i
	}
}

func consumer(in <-chan int, idx int) {
	for num := range in {
		fmt.Printf("-----消费者%dth，消费：%d\n", idx, num)
	}
}

func main() {
	product := make(chan int)
	rand.Seed(time.Now().UnixNano())
	go producer(product, 1) // 1 生产者
	go consumer(product, 1) // 3 个消费者
	for {
		time.Sleep(1000)
	}
}
