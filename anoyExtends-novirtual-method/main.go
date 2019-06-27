package main

import "fmt"

type Showable interface{
	ShowA()
	ShowB()
	}

	type Person struct{

	}

func (this *Person) ShowA() {
	fmt.Println("Person ShowA")
	//go 无法从父类 指向 子类实现 ,但是它可以实现子类覆盖父类 方法
	this.ShowB()
}


func (this *Person) ShowB() {
	fmt.Println("Person ShowB")
}

type Student struct{
		Person
}

func (this * Student) ShowB() {
	fmt.Println("Student ShowB")
}

// go的 多态是，通过接口，实现的完全多态； 基于接口的 行为完全替换
//说白了，就是 struct 基类是 无法作为 基类抽象的，有且只能接口类型充当 基类抽象
func main()  {

	s := Student{}
	//s.ShowA()

	var i Showable = &s
	i.ShowA()
	i.ShowB()



}
