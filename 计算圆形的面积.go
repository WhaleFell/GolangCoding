package main

import (
	"fmt"
	"math"
)

func main() {
	r := -2.0
	area, err := circleArea(r)
	if err != nil {
		// 打印 err 就是调用错误struct的Error()方法,返回字符串
		// fmt.Println(err)
		// 通过接口断言判断错误类型,获取错误类型中的属性
		if err, ok := err.(*areaError); ok {
			fmt.Println(err)
			fmt.Printf("半径是:%.2f\n", err.radius)
		} else {
			fmt.Println(err)
			fmt.Printf("其他错误类型:%s,%T\n", err, err)
		}
		return
	}
	fmt.Println("圆的的面积是:", area)
}

//1. 定义一个struct,表示错误类型
type areaError struct {
	msg    string
	radius float64
}

//2. 实现error接口计算实现Error()方法
// 需要传入结构体指针
func (e *areaError) Error() string {
	// 返回一个格式化的字符串
	return fmt.Sprintf("error: 半径,%.2f,%s", e.radius, e.msg)
}

// 3.定义一个求圆面积的函数
func circleArea(r float64) (float64, error) {
	if r < 0 {
		// 因为是指针用 & 取真实的值,(错误类型结构体)
		return 0, &areaError{"半径是非法的", r}
	}
	return math.Pi * r * r, nil
}
