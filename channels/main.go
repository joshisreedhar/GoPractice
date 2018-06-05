package main

import (
	"time"
)

func main() {

	chs := make(chan string)
	chr := make(chan string)
	go where(chs, "world")
	go greeting(chr, "hello")
	for i := 0; i < 2; i++ {
		select {
		case msg := <-chs:
			print(msg)
		case msg2 := <-chr:
			print(msg2)
		case <-time.After(1 * time.Second):
			print("timeout")
		default:
			print("No Activity")
		}

	}

}

func where(ch chan<- string, msg string) {
	time.Sleep(1 * time.Second)
	ch <- msg
}

func greeting(chr chan<- string, msg string) {
	time.Sleep(3 * time.Second)
	chr <- msg
}
