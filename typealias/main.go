package main

import "fmt"

type Integer int
type Integer2 =int


func main() {

	var a Integer2 =10
	var b Integer2 =20

	fmt.Println(a+b) //output: 30


	var c Integer =-20
	if c==a{ //认为是不同类型，编译报错
		fmt.Println("hello world...")
	}


}
