package main

import "fmt"

func main() {
	/*
		数据类型
			基本类型: int, fload, bool, string
			复合类型: array, slice, map, struct, pointer, function, channel(通道)
		数组:
			1. 储存一组相同数据类型的数据结构,定长的容器
				理解为容器, 储存一组数据,数组是定长的

		定义数组: var array_name [SIZE] variable_type

	*/
	// step1: 创建数组,定义数组的边界为4
	var arr1 [4]int
	// step2: 数组的访问
	// 不能超出数组的访问
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	fmt.Println(arr1[0])
	// 适用于 len(array/map/slice/string)
	fmt.Println("数组的长度:", len(arr1)) //容器中实际存储的数据量
	fmt.Println("数组的容量:", cap(arr1)) //最大存储
	//数组定长,长度=容量

	// 数组的其他创建方式
	var a [4]int   // 同: var a=[4] int
	fmt.Println(a) //默认全部 [0,0,0,0]

	var b = [4]int{1, 2, 3, 4}
	fmt.Println(b)

	var c = [5]int{1: 1, 3: 2} //传入下标创建数组
	fmt.Println(c)

	var e = [5]string{"rose", "hyy", "love"}
	fmt.Println(e)

	f := [...]int{1, 2, 3, 4} //不定长的数组
	fmt.Printf("%v %d %d", f, len(f), cap(f))

	g := [...]string{1: "hyy", 3: "黄颖怡"} //不定长的数组传入下标创建
	fmt.Printf("%v %d %d", g, len(g), cap(g))

}
