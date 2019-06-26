package main

import "fmt"

type People interface {
	Show()
}

type Student struct{}
func (stu *Student) Show() {

}

//可以看出 go的 *type指针类型，如果转化为 interface{} 接口类型，会有 装箱和拆箱的过程，里头还是包着一个nil
//记住 也只有， *type 指针类型才可以装箱。。。

func getPerson() People {
	var p *Student
	//fmt.Println("not nil")
	if p==nil{
		fmt.Println("hello world ..nil..")
	}else{
		fmt.Println("hello world ...not nil.")
	}
    //返回的过程中，装箱为interface
	return  p
}





//用了一个interface{} 来包装一个指针，甚至是null ptr 空指针，防止程序异常
func main() {
	if getPerson()==nil{ //已经被装箱了，所以判断不可能是nil, 但是真实效果上 ，里面且是nil
		fmt.Println("nil")
	}else{
		fmt.Println("not nil")
		fmt.Printf("%v",getPerson())
	}

}
