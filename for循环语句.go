package main

import (
	"fmt"
	"math"
)

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

	//1. 打印 58-23 的数字
	for i := 58; i >= 23; i-- {
		fmt.Println("打印58-23的数字:", i)
	}

	sum := 0
	//2. 求1-100的和
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1-100的和", sum)

	//3. 打印 1-100 内能被3整除,但是不能被5整除的数字,统计被打印的数字个数,每行打印5个
	count2 := 0 //计数器
	for i := 1; i <= 100; i++ {
		if (i%3 == 0) && (i%5 != 0) {
			// fmt.Println("数字:", i)
			fmt.Print(i, "\t")
			count2++ //5,10,15 换行
			if count2%5 == 0 {
				fmt.Println()
			}
		}
	}
	fmt.Println("\n", "count---->", count2)

	// 多层 for 循环
	for i := 0; i < 5; i++ {
		for i := 0; i < 5; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//打印乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d   ", j, i, i*j)
		}
		fmt.Println()
	}

	/*
		循环控制语句:
		break: 跳出循环,彻底结束循环
		continue: 跳出一次循环

		注意点: 多层循环嵌套,break和continue,默认结束这个循环,
		可以给循环其别名跳出哪个循环
		break 循环标签名.
	*/
	for i := 0; i < 10; i++ {
		if i == 2 {
			// break  终止这个 for 循环
			continue
		}
		fmt.Println(i)
	}

	// 给循环起别名.
out:
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d   ", j, i, i*j)
			if j == 2 {
				break out // 跳出最外层循环.
			}
		}
		fmt.Println()
	}
	/*
		水仙花数: 三位数: [100,999]
			每个位上的数字的立方和.刚好等于该数字本身.那么就叫水仙花数.
			e.g. 153
				1^3+5^3+3^3=153
	*/
	fmt.Println("\n------------------------------------")

	for a := 1; a <= 9; a++ {
		for b := 1; b <= 9; b++ {
			for c := 1; c <= 9; c++ {
				sum := a*100 + b*10 + c
				// val := a*a*a + b*b*b + c*c*c
				// math.Pow 幂运算
				val := math.Pow(float64(a), 3) + math.Pow(float64(b), 3) + math.Pow(float64(c), 3)
				if float64(sum) == val {
					fmt.Println("水仙花数抓住它了!", sum)
				}
			}
		}
	}

	/*
		打印 2-100 内的质数(只能被1和它本身整除)
	*/
	for i := 2; i <= 100; i++ {
		// 优化: flag 记录是否是质数
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
			break
		}
		if flag {
			fmt.Println("质数抓住它了!", i)
		}
	}

	/*
		goto 语句
			可以无条件地转移到过程中指定的行,通常结合判断语句使用
		label: statement 标签语法
	*/
	var a float64 = 10

LOOP:
	for a > 20 {
		if a == 15 {
			a += 1
			goto LOOP
		}
		fmt.Printf("a=%d\n", a)
	}

	// goto 集中处理错误
	fmt.Println("-----------------------")
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if j == 2 {
				goto breakHere
			}
		}
	}
	//手动返回,避免执行入标签
	return

breakHere:
	fmt.Println("done.....")
}
