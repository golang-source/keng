package main

import "fmt"

func main() {

	var c =0

	loop1:
		if (c==10){
			goto end
		}
		c++

	for i := 0; i < 10; i++ {
		innerflag:
		println(i)
		goto end
	}
	goto loop1
	// goto innerflag ;error 只能往外部跳，但是不能跨函数


end:

	fmt.Println("has exit c=",c)

}