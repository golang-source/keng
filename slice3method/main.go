package main

import "fmt"


//cap 的作用就是提前分配内存；通过提前申请内存，提升运行过程中的性能损耗
//length的就是控制数据元素写到 哪里了？ n_writed_pos,影响下次写入的位置

func main()  {
	//cap  and length  cap=length=5
	fmt.Println("param 2")
	var  s =make([]int,5)   //   s : [0,0,0,0,0]

    //when cap==length ,append-opt will append to last
	s = append(s,1,2,3,4,5)  //will 0 0 1 2 3
	fmt.Println(s)

	fmt.Println("'-----------param 0-----")
	//性能最差的方式就是 无参数方式; 动态方式，有几个分配几个空间
	var s2 =make([]int,0)   //equ   var s2 []int
	s2=append(s2,1,2,3,)
	fmt.Println(s2) //print : 1 2 3 0 0


	fmt.Println("'-----------param 3-----")

	//3个参数方式和两个参数 性能差不多， 无参数的方式只有 他们的 40%的性能
	var s3 = make([]int,0,5)

	s3 = append(s3,1,2,3,4,5)
	fmt.Println(s3)
}
