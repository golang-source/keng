package main

import "fmt"

func doMap(m map[int]string,indx int) (string,bool) { //多个参数得加括号，一个参数如果有名字也得加括号
	if  v,exist:=m[indx];exist{
		fmt.Println("has exist ....")
			return  v ,true
		}
		//return nil , false   // error string 没有nil值；
		// nil 只针对
		//1 系统引用类型： 如 map chan  slice
		//2 指针类型
		return "",false
}

func main() {
	m:=map[int]string{
			1:"wnewep",
			2:"stevejob",
	}

	v,b:=doMap(m,3)
	if !b {
		fmt.Printf("not exist index=%d v:",3,v)
	}
}