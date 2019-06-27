package main

import (
	"fmt"
)

func doWrite(b chan int ,writeChan chan string)  {
	var i = 0
	for {
		c:=<-b
		if c!=0{
			s:=fmt.Sprintf("id=%d helloworld",i)
			writeChan<-s
			fmt.Println(s)
			i++
		}
	}

}

func main() {
	var bucket = make(chan int)
	var writeChan =make(chan string,100)

	go doWrite(bucket,writeChan)

	//controll speed to decide write speed  频率
	for i:=0;i<101;i++{
		bucket<-i
		//time.Sleep(time.Millisecond*500)
	}

	defer close(bucket)
	defer close(writeChan)

}
