package main

import "fmt"

//struct 可以比较，但是有条件
//1. 不能有 ms 这种不可以比较的元素
//2、如果不存在ms元素，那么对  ，属性的个数和顺序必须一致，否则不可比较，会编译期错误
// 如果 没有ms 元素，顺序 个数 名字 都一样，那就可以比较
func main() {
	sn1 := struct {
		age  int
		name string
		c chan int

	}{age: 11, name: "qq",c:make(chan int)}
	sn2 := struct {
		age  int
		name string
		c chan int
	}{age: 11, name: "qq",c:make(chan int)}

	if sn1 == sn2 {  //因为c 通道不一样
	fmt.Println("sn1 == sn2")
}

/*
sm1 := struct {
age int
m   map[string]string
}{age: 11, m: map[string]string{"a": "1"}}
sm2 := struct {
age int
m   map[string]string
}{age: 11, m: map[string]string{"a": "1"}}

if sm1 == sm2 {
fmt.Println("sm1 == sm2")
}
*/



}