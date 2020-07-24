package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c := make(chan int) // チャネルを作成

	callNum := 5

	for i := 0; i < callNum; i++ {
		go sleepyGopher(i, c) // goroutine
	}

	for i := 0; i < callNum; i++ {
		gopherID := <-c // チャネルで値を受信
		fmt.Println("gopher", gopherID, "はスリープを終えました。")
	}

	//time.Sleep(4 * time.Second)
}

func sleepyGopher(id int, c chan int) {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(4)
	time.Sleep(time.Duration(randomNum) * time.Second)
	fmt.Println("gopher", id, "は、", randomNum, "秒寝ます。 ... snore ...")
	c <- id // チャネルに値を送信
}
