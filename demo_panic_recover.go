package main

import "fmt"

func main() {
	// 处理 panic 的 defer 匿名函数
	defer func() {
		// reconver() 接受 panic() 的信息
		if msg := recover(); msg != nil {
			fmt.Println(msg, "程序恢复啦。。")
		}
	}()
	// funB的panic传递到此处,引起panic时需要执行完所有defer主函数才会报错
	funB()

}

func myprint(s string) {
	fmt.Println(s)
}

func funB() { //外围函数
	// 匿名函数,处理pianic的recover()也可以放在main函数中
	// defer是逆序执行的
	defer func() {
		// reconver() 接受 panic() 的信息
		if msg := recover(); msg != nil {
			fmt.Println(msg, "程序恢复啦。。")
		}
	}()

	fmt.Println("我是函数funB()..")
	defer myprint("defer funB()...1.....")
	for i := 1; i <= 10; i++ {
		fmt.Println("i:", i)
		if i == 5 {
			//让程序中断
			panic("funB函数,恐慌了")
		}
	}
	//当外围函数的代码中发生了运行恐慌，只有其中所有的已经`defer`的函数全部都执行完毕后，
	//该运行恐慌才会真正被扩展至调用处。
	defer myprint("defer funB():2.....")
}
