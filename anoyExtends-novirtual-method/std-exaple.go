package main

import (
	"fmt"
	"strings"
)

type IAnimal interface {
	 Eat()
	 Say()
}

type Cat struct {
   Name string
}

type Dog struct {


}

func (this * Cat) Eat()  {
	fmt.Println("cat eat")
}
func (this * Cat) Say()  {
	fmt.Println("cat say ")
	//fmt.Println("cat say ,name:",this.Name)
}


func (this * Dog) Eat()  {
	fmt.Println("Dog eat")
}
func (this * Dog) Say()  {
	fmt.Println("Dog eat")
}

func factory_getAnimal(t string) IAnimal {
			if strings.ToLower(t)=="dog"{
				return &Dog{}
			}else if strings.ToLower(t)=="cat"{
				return &Cat{}
		}
		//var defaultAmi *Cat  // here is nil  ，这里是nil
	//but when  as interface to  box defaultAmi to be not-nil interface
	//也就说，虽然 defaultAmi 是 nil 指针，但是如果 转化为 interface，系统会将其box 装箱
	//用 not nil interface 类型 来表达 nil的值

	//return  defaultAmi

		return nil
}

func doExecute(i IAnimal)  {
	i.Say()
	i.Eat()
}

func main() {
	var i IAnimal = factory_getAnimal("dog")
	doExecute(i)

	i = factory_getAnimal("cat")
	doExecute(i)

	i =factory_getAnimal("bird")  //nil interface ,被包装过指针类型，是可以访问函数的，但是无法访问属性，否则崩溃
	if i==nil{
	   fmt.Println(" return nil ")
	}else{
		fmt.Printf("%p \n",i)
		fmt.Println(" return not nil ")


	}

	doExecute(i)


}