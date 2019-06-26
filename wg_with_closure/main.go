package main

import (
	"runtime"
	"sync"
	"fmt"
)

func main() {
	runtime.GOMAXPROCS(1)

	wg :=sync.WaitGroup{}
	wg.Add(20)


	//一句话：闭包延迟执行
	for i:=0 ;i<10 ;i++{
		go func() {
			fmt.Println("A:",i)   //闭包延迟执行函数，所以退出for 后，一个个执行函数时候，i已经等于10
			wg.Done()
		}()

	}

	for i:=0;i<10;i++{
		go func(n int ) {
				fmt.Println("B:",n)  //通过外部传入 ，避免闭包 污染
				wg.Done()
		}(i)

	}

	wg.Wait()
}
