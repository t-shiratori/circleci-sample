package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello go")

	c := make(chan int) // チャネルを作成

	callNum := 5

	for i := 0; i < callNum; i++ {
		go sleepyGopher(i, c) // goroutine
	}

	for i := 0; i < callNum; i++ {
		gopherID := <-c // チャネルで値を受信
		fmt.Println("gopher", gopherID, "スリープを終えました")
	}

	//time.Sleep(4 * time.Second)
}

func sleepyGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...", id, "snore ...")
	c <- id // チャネルに値を送信
}
