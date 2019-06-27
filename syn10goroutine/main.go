package main

import (
	"time"
	"fmt"
)


const size=10

func main()  {


	var waitChan =make (chan int,size)

	for i:=0;i<size;i++ {

		go func(signal chan int) {
			for j:=0;j<10;j++{
				fmt.Println("index:",j)
				time.Sleep(time.Second)
			}
			signal<-1
			fmt.Printf("signal=%v ...\n",signal)
		}(waitChan)

	}


	for{
		time.Sleep(time.Second)
		fmt.Printf("waitSlice=%v ...\n",waitChan)
		if len(waitChan)==size{
			break
		}
	}



}
