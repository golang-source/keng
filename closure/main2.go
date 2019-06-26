package main

import "fmt"

//由于延迟闭包，最后离开for i=2, 所以执行 闭包函数的时候，i那个时候 就是2，
func getFuns() []func() {
	//无需参数，三种切片初始化方式里面性能最差的方式;
	// 参数2 make([]type,cap)   容量和长度都不变情况
	// 参数3 make([]byte,len,cap)// 容量不变，长度会变
	//
	var funs []func()

	for i:=0;i<2;i++{
		funs= append(funs, func() {
			fmt.Println(&i,i)    //because of closure and closure will delay execute , so at last break for i=2 ,
		})
	}
	return funs
}


func main()  {

	var fs =getFuns()
	for _,f:=range fs{
		f()
	}

}
/*output
0xc042008098 2
0xc042008098 2*/
