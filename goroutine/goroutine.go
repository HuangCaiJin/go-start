package main

import (
	"fmt"
	"time"
)

func main() {
	go runtimes(10)
	for i := 1; i <= 10; i++ {
		fmt.Println("main",i,"main-剩余",10 - i)
		time.Sleep(2 * time.Second)
	}
}

func runtimes(times int) int {
	for i := 1;i <= times; i++ {
		fmt.Println("runtimes",i,"剩余次数",times - i)
		time.Sleep(time.Second)
	}
	return times
}