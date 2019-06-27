package main

import (
	"sync"
	"fmt"
)

//下面的迭代会有什么问题？

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	 //ch := make(chan interface{}) // 解除注释看看！
	 fmt.Println("len=",len(set.s))

	ch := make(chan interface{},len(set.s))
	go func() {
		set.RLock()

		for elem,value := range set.s {
			ch <- elem
			println("Iter:",elem,value)
		}
		close(ch)
		set.RUnlock()

	}()
	return ch
}

//主线程执行时间越持久，其它 就有被调度机会
func main()  {
	var sarr []interface{}
	for i:=0;i<10;i++{
		sarr =append(sarr,i)
	}
	th:=threadSafeSet{
		s:sarr,
	}
		//对通道的读写 会产生挂起
		v:= th.Iter()  //主线程挂起等待，那么 go func 就会有执行机会，到底执行多少个循环 或者指令周期，

	for i:=0;i<10;i++{
		s:=<-v
		fmt.Printf("%v \n",s)
	}


}