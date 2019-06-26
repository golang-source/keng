package main

import "fmt"

//panic 会中断后面的代码执行流程
//但是如果后面有代码再次执行 panic ，那么它会覆盖之前的panic

func main() {

	defer func() {
		if e:=recover();e!=nil{
			fmt.Println(e)
		}
	}()

	defer func() {
		panic("panic2")
	}()

	fmt.Println("helloowrld")
	panic("panic1 ")
}
