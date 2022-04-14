package main

import "fmt"

func main() {
	/*
		for 循环
		语法:
		for 表达式1(init);表达式2(condition);表达式3(post){
			循环体
		}
		过程:
		初始化语句(init)->检查条件(condition)->true->循环体->
		post->重新检查条件(condition)->true->
		循环体->post->重新检查条件(condition)->true->
		循环体

	*/

	count := 10
	// i 的作用域只在 for 循环当中
	for i := 0; i <= count; i++ {
		fmt.Println("for 循环")
	}

	/*
		for 循环的其他写法
		for 的三个部分都是可选的

		for condition{ }  相当于: while(条件)

		省略了表达式2(condition)相当于直接作用在 true 上
		同时省略3个表达式(相当于无限循环) while True:
		for ; ; ;{

		}

	*/
	// i 的作用域不同
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	/* for{
		fmt.Println("无限循环")
	} */
}
