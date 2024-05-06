package main
import (
	"fmt"
	"rand"
	"time"
)

var intChan chan int


func main() {

	fmt.Println("123")

}



func readData() {

}


func writeData(intChan chan int) {
	rand.Seed(time.Now().UnixNano()))

	for i := 1; i < 50; i++ {
		tempInt := rand.Intn(4) + 10
		intChan <- tempInt
		
	}
	close(intChan)
}