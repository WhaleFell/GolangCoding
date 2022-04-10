package main

import "fmt"

func a() {

}

func main() {
	var name string = "tom"
	age := 20
	b := true
	// 格式化输出
	// %T 类型
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", b)

	//指针类型
	a := 100
	p := &a             //取a的地址 p是指针类型
	fmt.Printf("%T", p) //*int

	//数组类型
	/* b := [2]int{1, 2}
	fmt.Printf("%T\n", b) */

	//切片类型 (动态数组)
	/* a := []int{1, 3, 4}
	fmt.Printf("%T", a) */

	//函数
	/*
		fmt.Printf("%T", a) */
}
