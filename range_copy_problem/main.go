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


func main() {
	m:=pauseMap()
	for k,v:=range  m{
		fmt.Printf("name:%s ,v=%v",k,v)
	}
}
