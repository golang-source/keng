package main



// ins.type() 只能用于接口类型interface{} 其它指针类型要转化为interface{}才可以操作，所以编译错误
func main() {
	i := GetPerson()

	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}

}

type Person struct{


}

func GetPerson() *Person{
	return &Person{}
}

func GetValue() int {
	return 1
}
