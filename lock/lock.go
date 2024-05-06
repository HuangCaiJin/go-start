package main

import (
	"fmt"
	"time"
	"sync"
)

var (
	testMap = make(map[int]int)
	lock sync.Mutex
)



func testNum(num int) {
	lock.Lock()
	res := 1
	for i := 1; i <= num; i++ {
        res *= i
    }
	testMap[num] = res

	lock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 20; i++ {
		go testNum(i)
    }	
	// 协程需要再main结束前执行完毕
	// time.Sleep(time.Second * 8)
	lock.Lock()
	for key,value := range testMap {
		fmt.Printf("数字是%v 对应的阶乘是 %v \n",key,value)
	}
	lock.Unlock()
	end := time.Since(start)
	fmt.Println("耗时：",end)

}