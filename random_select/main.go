package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(1) //分析goroutine

	var intChan =make(chan int,1)
	var strChan =make(chan string,1)

	intChan<-1
	strChan<- "hell world"

	//加入 runtime.GOMAXPROCS(1) 是 1 个核心，那 select 的case  随机 选择“一个”的机会更大；
	// 如果是多个核心，在随机一个的基础上，出现两个一起执行机会更大;  但是不管单核还是多核，select 选择case 执行就是随机性
	select {
	//select keng:select的随机选择性
	//其实select 的每个case都是并发结构 并行结构
		    case  v:=<-intChan:
			fmt.Println(v)

			case s:=<-strChan:
				fmt.Println(s)
				panic(s)

	}
	//主线程不可能退出 ，永远得不到执行机会 ，就会报 deadlock
}
