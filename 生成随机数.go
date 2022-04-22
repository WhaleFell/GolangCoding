package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		生成随机数 random:
			伪随机数,根据一定的算法公式算出来的
		math/rand
	*/
	// 需要给一个种子数来生成随机数,但是默认提供了一个种子数
	// 内置种子数不变,所以产生的随机数也不变
	num1 := rand.Int()
	fmt.Println(num1)

	for i := 0; i < 10; i++ {
		num := rand.Intn(10) //产生随机范围的随机整数 [0,10)
		fmt.Println(num)
	}

	rand.Seed(1) //设置种子数,种子数改变生成的随机数也改变
	num2 := rand.Intn(1000)
	fmt.Println("---->", num2)

	//通过时间来改变种子数,因为时间是不断改变的.
	t1 := time.Now()            //type: time.time对象
	timeStamp1 := t1.Unix()     // 时间戳 秒
	timeStamp2 := t1.UnixNano() //时间戳 纳秒
	fmt.Printf("%v %v", timeStamp1, timeStamp2)

	// 生成真正的随机数
	// step1: 设置种子数,可以设置成时间戳
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		// 调用生成随机数的函数
		fmt.Println("--->", rand.Intn(100))
	}
}
