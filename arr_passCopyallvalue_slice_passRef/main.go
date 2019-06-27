package main

import "fmt"

//copy arr all elements
func changeArr(a [3]int)  {
	 a[0]=10
}

func changeSlice(a []int)  {
	a[0]=333
}

func main() {
	arr:= [3]int{1,2,3}
	changeArr(arr)

	fmt.Println(arr)
	changeSlice(arr[:]) //slice就是一个array 的 boxing
	//go 中 interface 和指针之间有boxing ，slice 和array之间有boxing

	fmt.Println(arr)

}
