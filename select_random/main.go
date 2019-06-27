package main

import (
	"runtime"
	"fmt"
	"time"
)

//基于超时机制的 select 读取全部数据方法
//0 读取全部数据，采用 for{select}
//1 采用了for{},导致一直循环堵塞，主协程，得不到调度(由于main主协程 进行了io读写 有堵塞)，所以就会deadlock错误
func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"

    var tt  = time.NewTimer(time.Second)
	for{
		tt.Reset(time.Second)

		select {
		case value := <-int_chan:
			fmt.Println("int_chan:",value)
		case value := <-string_chan:
			fmt.Println("string_chan:",value)
			//select
	/*	case t:=<-time.After(time.Second):
			fmt.Println("out of program,",t)
			goto end*/
		case t:= <-tt.C:  //opt version
		fmt.Println("out of program,",t)
		goto end
		}
	}

	end:
		tt.Stop()
		fmt.Println("will exit")
}


