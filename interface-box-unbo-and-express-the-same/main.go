package main

import "fmt"

type Human interface{
	Say(string)
}

type Student struct{
	name string

}

func (this * Student) Say(msg string) {
	fmt.Println(msg,this.name)
}

func main()  {

	var h Human
	// from right value to left interface{} ,so boxing
	h = &Student{name:"wenweiping"}
	//when call ，unboxing to call inner-real Student ptr
	h.Say("hello world")

}
