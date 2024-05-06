package main

import "fmt"

//创建一个接口
type tank interface {

    // 方法
    Tarea() float64
    Volume() float64
}

type task interface {

    // 方法
    Sum() float64
}

type myvalue struct {
    radius float64
    height float64
}

//实现方法
//桶的（Tank）接口
func (m myvalue) Tarea() float64 {

    return 2*m.radius*m.height + 2*3.14*m.radius*m.radius
}

func (m myvalue) Volume() float64 {

    return 3.14 * m.radius * m.radius * m.height
}


func (m myvalue) Sum() float64 {
	return m.radius + m.height
}

func main() {

    // 访问使用桶的接口
    var t tank
    t = myvalue{10, 14}
    fmt.Println("桶的面积 :", t.Tarea())
    fmt.Println("桶的容量:", t.Volume())

	var t1 task
	t1 = myvalue{20,15}
	fmt.Println(t1.Sum())
	
	
}