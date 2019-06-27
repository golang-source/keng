package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func pauseMap()map[string]*student {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus { //由于共享 一个变量 stu， range 是copy 数值的，导致
		m[stu.Name] = &stu   //  &stu 一致引用了一个变量，最后这个变量的值 就是最后一个元素，就是key 不一样而已
	}
	return m
}

//for range 就是容易犯，维持最后一个 元素的值问题: 解决中间变量就行
//闭包常见问题：go func 和 defer  func ，panic 都具有延时性， 就是延迟求值，导致被引用的值 成为了最后一个元素的值
func Resovle()map[string]*student {

	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
	    s:=stu  //copy 中间变量解决问题
		m[s.Name] = &s
	}
	return m
	
}

func main() {
	m:=Resovle()
	for k,v:=range  m{
		fmt.Printf("name:%s ,v=%v",k,v)
	}
}
