package main

import "fmt"

func main() {
	//list := make([]int,0)  //ok 1
	list:=new ([]int)   //ok 2
	s := append(*list, 1)  //注意: *list
	fmt.Println(s)
}
