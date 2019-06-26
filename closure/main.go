package main

import "fmt"

func getFun(x int) (func(),func()) {
	//下面函数闭包，引用了共同外变量，主动相互影响
	return func() {
			fmt.Println(x) //100
			x+=1    //x =101
	},
	func() {
			fmt.Println(x)
		}
}



func main() {

	a,b :=getFun(100)
	//because of a fun , run first
	a()   //x =100 run after x turn x=101
	b()   //then x =101

}
