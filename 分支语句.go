package main

import "fmt"

func main() {
	/*
		条件语句:if
		语法格式:
			if 条件表达式{

			}
		注意点:
			1.if 后的{,一定是和 if 条件写在同一行的
			2.
	*/
	num := 16
	if num > 10 {
		fmt.Println("大于10")
	}

	score := 0
	fmt.Println("请输入你的成绩:")
	fmt.Scanln(&score) //通过键盘读取 内存地址赋值给变量

	if score >= 60 {
		fmt.Println("score:", score, "及格")
	} else {
		fmt.Println("成绩不及格", score)
	}

	var sex string = "男"
	if sex == "男" {
		fmt.Println("去男厕所.")
	} else if sex == "女" {
		fmt.Println("去女厕所.")
	} else {
		fmt.Println("去第三卫生间.")
	}

	// if 语句的嵌套
	if sex == "男" {
		fmt.Println("去男厕所.")
	} else {
		// 嵌套if
		fmt.Println("去第三卫生间.")
	}

	/*
		if 语句的其他写法
		if 初始化语句; 条件{

		}
	*/
	// num 的作用域仅仅限于 if 块里面
	// 随着 if 块的结束而销毁
	// num2 := 5  // 作用域是整个main函数

	if num := 4; num > 0 {
		fmt.Println("正数。。", num)
	} else if num < 0 {
		fmt.Println("负数。。", num)
	}

	/*
		switch 语句
			实现选择结构
			swith 变量名{
				case val1:
				分支1
				case val2:
				分支2
				case val3:
				分支3
				default:
				最后一个分支.
			}
		注意事项：
			1. switch 可以作用在其他类型上，case后的数值必须和switch作用的变量类型一致.
			2. case 是无序的
			3. case 后的数值是唯一的
			4. default: 可选的操作
	*/
	num3 := 3
	// case 一旦匹配成功 整个 switch 分支结束
	switch num3 {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("not found")
	}

	// 省略 switch 后的变量。相当于直接作用在 true 上
	switch true {
	case true:
		fmt.Println("true...")
	case false:
		fmt.Println("false....")
	}

	/*
		case 后可以同时跟随多个数值
		switch 变量{
		case val1, val2, val3:
			.....
		case val4, val5:
			.....
		}
	*/

	/*
		switch 中的 break 和 fallthrough 语句

		break: 可以使用在 switch 中也可也使用在 for 循环中.
			强制结束 case 语句,从而结束 switch 分支
		fallthrough: 贯通后续的 case 用于switch的穿透
			应该位于 case 块的最后一行
	*/
	n := 2
	switch n {
	case 1:
		fmt.Println("我是熊大")
		fmt.Println("我是熊大2")
		fmt.Println("我是熊大3")
	case 2:
		fmt.Println("我是熊二")
		fmt.Println("我是熊二2")
		break //强制结束 case 块
		fmt.Println("我是熊二3")
	case 3:
		fmt.Println("我是光头强")
		fmt.Println("我是光头强2")
		fmt.Println("我是光头强3")
	}

	m := 2
	switch m {
	case 1:
		fmt.Println("我是熊大")
	case 2:
		fmt.Println("我是熊二")
		fallthrough //紧接着执行下一个 case(switch的穿透)
	case 3:
		fmt.Println("我是光头强")
	}

	fmt.Println("main..over..")
}
