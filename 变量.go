package main

import "fmt"

func getNameAndAge() (name string, age int) {
	return "hyy", 30
}

func main() {

	//调用函数,匿名变量
	name, _ := getNameAndAge()
	fmt.Printf("name: %v\n", name)

	var name1 string
	fmt.Printf("name: %v\n", name1)

	// 批量声明
	var (
		name string
		age  int
		b    bool
	)

	name = "hyy"
	age = 14
	b = false

	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("b: %v\n", b)

	//变量初始化
	var title string = "lovehyy"
	var site string = "http://baidu.com"
	var agge int = 30

	// 类型推断
	var name = "hyy"
	var age = 20
	var b = true
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("b: %v\n", b)

	//批量初始化
	var name, age, b = "tom", 20, true
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("b: %v\n", b)

	//短变量声明
	name := "hyy"
	age := 51
	b := true
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("b: %v\n", b)

}
