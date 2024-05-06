package main

import (
	"fmt"
	// "os"
	// "unicode/utf8"
	// "strings"
)


func main() {
	// fmt.Println("this is a test message")
	// fmt.Println("1 + 1 =", 1 + 1)
	// var Pname = "(cainiaojc.com)" 
    // var Lname = "Go 语言" 
    // var Cname = "关键字"
      
    // fmt.Printf("站点域名: %s", Pname) 
    // fmt.Printf("\n语言名称: %s", Lname) 
    // fmt.Printf("\n章节名称: %s", Cname) 

	// var x int8 = -100
	/* x := -200
	fmt.Println(x)

	a := 2.1
	b := 5.6

	c := b - a
	fmt.Println("结果: %f", c)
	fmt.Printf("\nc的类型是 : %T \n", c)   


	var d complex128 = complex(4,6)
	var e complex64 = complex(8,9)
	fmt.Println(d)
	fmt.Println(e)


	str1 := "cainiaojc"
    str2 := "cainiaojc"
    str3 := "cainiaojc"
    result1 := str1 == str2
    result2 := str1 == str3

    //打印结果
    fmt.Println(result1)
    fmt.Println(result2)

    //显示result1和result2的类型
    fmt.Printf("result1 的类型是 %T ， "+"result2的类型是 %T", result1, result2)


	str := "123"
	str = "234"
	//显示字符串的长度
	fmt.Printf("\n字符串的长度:%d", len(str)) 
     
	//显示字符串 
	fmt.Printf("\n字符串是: %s", str) 
	  
	// 显示str变量的类型
	fmt.Printf("\nstr的类型是: %T", str) 


	var myvariable1 = 20
    var myvariable2 = "cainiaojc"
    var myvariable3 = 34.80

    // Display the value and the
    // type of the variables
    fmt.Printf("myvariable1的值是 : %d\n", myvariable1)

    fmt.Printf("myvariable1的类型是 : %T\n", myvariable1)

    fmt.Printf("myvariable2的值是 : %s\n", myvariable2)

    fmt.Printf("myvariable2的类型是 : %T\n", myvariable2)

    fmt.Printf("myvariable3的值是 : %f\n", myvariable3)

    fmt.Printf("myvariable3的类型是 : %T\n", myvariable3)

	i,j := os.Open("./test.txt")

	fmt.Println(i,j == nil)

	const A= "2"
    var B = "1"
	fmt.Println(B < A) 


	a1 := 94
     
	//使用运算符（＆）和
	//指针间接（*）运算符
	b1 := &a1  
	fmt.Println(b1)
	fmt.Println(*b1)  
	*b1 = 67 
	fmt.Println(a1) 

	if(a1 > 100) {
		fmt.Println("a1 > 100")
	}else {
		fmt.Println("a1 <= 100")
	}
	
	for i := 0; i < 10; i++ {
		fmt.Println("i:",i)
	}

	s := 0
	for s < 20 {
		s++
		fmt.Println("s:",s)
	}

	rvariable:= []string{"GFG", "Geeks", "cainiaojc"} 

	for i, j:= range rvariable { 
		fmt.Println(i, j)  
	} 

	name := "huang cai jin"
	for c,s := range name {
		bytes := []byte{0, 0, 0, 0}
		nbs := utf8.EncodeRune(bytes, s)
		utf8Str := string(bytes[:nbs])
		fmt.Println(c,utf8Str)
	}

	// 假设我们有一个Unicode代码点
    codePoint := '你'
 
    // 将Unicode代码点转换为UTF-8字节序列
    bytes := []byte{0, 0, 0, 0} // 初始化一个足够长的字节切片来存储UTF-8编码
	fmt.Println(bytes)
	fmt.Println(rune(codePoint))
    numBytes := utf8.EncodeRune(bytes, rune(codePoint))
	fmt.Println(numBytes)
	fmt.Println(bytes[:numBytes])
    // 将字节切片转换为字符串并输出
    utf8Str := string(bytes[:numBytes])
    fmt.Println(utf8Str) // 输出UTF-8编码的字符


	mmap := map[int]string{
        22: "Geeks",
        33: "GFG",
        44: "cainiaojc",
    }
    for key, value := range mmap {
        fmt.Println(key, value)
    }


	chnl := make(chan string)
    go func() {
        chnl <- "abc"
        chnl <- "efg"
        chnl <- "hij"
        chnl <- "kln"
        close(chnl)
    }()
    for i := range chnl {
        fmt.Println(i)
    }
	fmt.Println("%t",codePoint)

	res4 := strings.IndexByte("cainiaojc, geeks", 'g')
	fmt.Println("结果 4: ", res4) */


	myslice := []string{"Geeks", "Geeks", 
                    "gfg", "GFG", "for"} 
      
    fmt.Println("Slice: ", myslice) 
  
    //使用比较运算符
    result1 := "GFG" > "Geeks"
    fmt.Println("Result 1: ", result1) 
  
    result2 := "GFG" < "Geeks"
    fmt.Println("Result 2: ", result2) 
  
    result3 := "Geeks" >= "for"
    fmt.Println("Result 3: ", result3) 
  
    result4 := "Geeks" <= "for"
    fmt.Println("Result 4: ", result4) 
  
    result5 := "Geeks" == "Geeks"
    fmt.Println("Result 5: ", result5) 
  
    result6 := "Geeks" != "for"
    fmt.Println("Result 6: ", result6) 



	//定义一个正常变量
    var x int = 5748

    //指针声明
    var p *int

    //初始化指针
    p = &x

    //显示结果
    fmt.Println("存储在x中的值 = ", x)
    fmt.Println("x的内存地址 = ", &x)
    fmt.Println("存储在变量p中的值 = ", p)
}