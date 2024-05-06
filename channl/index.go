package main

import (
	"fmt"
	"sync"
)

var intChan chan int
var lock sync.Mutex

func main() {
	intChan = make(chan int,10)
	intChan <- 5
	intChan <- 15
	fmt.Println(intChan)
	fmt.Printf("intChan的值是%v,地址是%v \n",<-intChan,&intChan)
	fmt.Printf("intChan的大小是%v,容量是%v \n",len(intChan),cap(intChan))

	fmt.Printf("intChan的值是%v,地址是%v \n",<-intChan,&intChan)
	fmt.Printf("intChan的大小是%v,容量是%v \n",len(intChan),cap(intChan))


	strChan := make(chan string,20)
	strChan <- "abc"
	strChan <- "def"

	fmt.Printf("strChan的值是%v,地址是%v \n",<-strChan,&strChan)
	fmt.Printf("strChan的大小是%v,容量是%v \n",len(strChan),cap(strChan))

	fmt.Printf("strChan的值是%v,地址是%v \n",<-strChan,&strChan)
	fmt.Printf("strChan的大小是%v,容量是%v \n",len(strChan),cap(strChan))


	mapChan := make(chan map[int]string,10)

	mapOne := make(map[int]string)
	mapOne[0] = "A-0"
	mapOne[1] = "A-1"

	mapChan <- mapOne
	
	mapTwo := make(map[int]string)
	mapTwo[0] = "B-0"
	mapTwo[1] = "B-1"

	mapChan <- mapTwo

	fmt.Printf("mapChan的值是%v,%v \n",<-mapChan,<-mapChan)
	fmt.Printf("mapChan的大小是%v,容量是%v \n",len(mapChan),cap(mapChan))


	allChan := make(chan interface{},10)
	allChan <- dog{name:"0001",color:"red"}
	allChan <- mapOne
	allChan <- "abc"
	allChan <- 123
	// fmt.Printf("allChan的值是%v,%v,%v,%v \n",<-allChan,<-allChan,<-allChan,<-allChan)

	// 类型断言 获取类型方法
	dogInfo := (<-allChan).(dog)
	// 类型 main.dog
	fmt.Printf("%T \n",dogInfo)

	fmt.Printf("%v \n",dogInfo.color)
	// for 循环管道取值前 需要close管道,关闭后不能在往管道里写入
	close(allChan)
	/* for val := range allChan {
		fmt.Println(val)
	}

	fmt.Printf("allChan的值是%v\n",<-allChan) */

	for {
		val,ok := <-allChan
		if !ok {
			break
		}
		fmt.Println(val)
	}
}


type dog struct {
	name string
	color string
}