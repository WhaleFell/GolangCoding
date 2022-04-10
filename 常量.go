package main

import "fmt"

func main() {
	const PI float64 = 3.14
	const PI2 = 3.1415 //可以省略类型

	// 批量声明
	const (
		width  = 100
		height = 200
	)

	//多重赋值
	const i, j = 1, 2
	const a, b, c = 1, 2, "foo"

	// 使用 `_` 跳过某些值
	const (
		a1 = iota //0
		_         //1
		a3 = iota //2
	)

	// `iota` 声明中间插队
	const (
		a1 = iota //0
		a2 = 100  //100
		a3 = iota //1
	)

	const (
		b1 = iota
		b2 = iota
		b3 = iota
	)

	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", b2)
	fmt.Printf("a3: %v\n", a3)

}
