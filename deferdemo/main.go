package main

import (
	"fmt"
	"time"
)



func main() {

	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
	}()

	defer func() {
		fmt.Println("3")
	}()

	defer func() {

		if e:=recover();e!=nil{
			fmt.Println("catch ",e)
		}else{
			fmt.Println("there is no error")
		}

	}()
	defer func() {
		time.Sleep(time.Second*10)
		if e:=recover();e!=nil{
			fmt.Println("catch2 ",e)
		}

	}()

	fmt.Println("hello world...")

	panic("触发异常1")
	panic("触发异常2") //第一个异常就会中断后面的代码
	panic("触发异常3")
}

/****
	hello world...
	3
	2
	1
/
 */
