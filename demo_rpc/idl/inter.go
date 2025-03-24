package main

import (
	"fmt"
	"strconv"
)

// 接口定义 & 实现
// Message 只是个声明 struct 来实现 interface?
// 动态调度方法的唯一方法是通过接口。结构或任何其他具体类型上的方法始终以静态方式解析
type Message interface {
	consume() string
	product() string
}

// ProductA 只能有声明属性字段（函数类型也是属性对吧）实现要分开
type ProductA struct {
}

func (p ProductA) consume() string {
	//TODO implement me
	panic("give a product A")
}

func (p ProductA) product() string {
	//TODO implement me
	panic("consume an product A")
}

// 方法实现
type upstream struct {
	name string
	age  int
}

func (u *upstream) product() Message {
	return ProductA{}
}

// 异常 - 实现了Error接口就是异常
type MyError struct {
	message string
	name    string
	msgCode int
}

func (e *MyError) Error() string {
	return e.name + ": " + strconv.Itoa(e.msgCode)
}

// defer 使用
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main1() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20

}

// channel 使用
func useChan() {

	//select 自动监听上下文中的ch
	ch := make(chan int)

	go func() {
		for {
			ch <- 29
		}
	}()
	go func() {
		for {
			ch <- 19
		}
	}()

	for {
		select {
		case <-ch:
			fmt.Println("done1")
		case <-ch:
			fmt.Println("done2")
		}
	}

}

func main() {
	useChan()
}
