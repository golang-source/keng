package main

import "fmt"

func deferCall()  {

	defer func() {
		if e:=recover();e!=nil{
			fmt.Println(e)
		}
	}()

	defer func() {
		panic("内部 ")
	}()

}

func funExecuteWithRecover(fn func())  {
	defer func() {
		if e:=recover();e!=nil{
			fmt.Println(e)
		}
	}()
	fn()
}


func main() {
	deferCall()
	//panic("helloworld")

}
