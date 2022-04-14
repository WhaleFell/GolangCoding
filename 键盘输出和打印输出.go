package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		输入和输出:
			fmt包: 处理输入和输出

			输出:
				Print() //打印
				Printf() //格式化打印
				PrintLn() //打印之后换行

			格式化打印占位符：
				%v,原样输出
				%T，打印类型
				%t,bool类型
				%s,字符串
				%f,浮点
				%d,10进制的整数
				%b,2进制的整数
				%o,8进制
				%x,%X，16进制
					%x：0-9，a-f
					%X：0-9，A-F
				%c,打印字符
				%p,打印地址

			bufio包读取
	*/

	a := 100           //int
	b := 3.14          //float64
	c := true          // bool
	d := "Hello World" //string
	e := `Ruby`        //string
	f := 'A'
	fmt.Printf("%T,%b\n", a, a)
	fmt.Printf("%T,%f\n", b, b)
	fmt.Printf("%T,%t\n", c, c)
	fmt.Printf("%T,%s\n", d, d)
	fmt.Printf("%T,%s\n", e, e)
	fmt.Printf("%T,%d,%c\n", f, f, f)
	fmt.Println("-----------------------")
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)

	var x int
	var y float64
	fmt.Println("请输入一个整数一个浮点类型:")
	fmt.Scanln(&x, &y) // 读取键盘的输入,通过操作内存地址,赋值给x和y 阻塞式函数
	fmt.Printf("a 的数值:%d,b的数值:%f\n", x, y)

	// 格式化输入
	fmt.Scanf("%d,%f", &x, &y) //输入 10,12.1
	fmt.Printf("x:%d,y:%f\n", x, y)

	// bufio包读取
	fmt.Println("--------------------------")
	fmt.Println("请输入一个字符串:")
	reader := bufio.NewReader(os.Stdin) // 从标准输出中读取
	s1, _ := reader.ReadString('\n')    //读到换行符号结束.读字符串
	fmt.Println("读到的数据", s1)
}
