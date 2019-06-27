package  main

import "fmt"

const (
	//iota 索引占位符，打断会重复
	x = iota  //after auto incr
	y
	z = "zz"
	k
	p = iota
)

func main()  {
	 fmt.Println(x,y,z,k,p)
}
