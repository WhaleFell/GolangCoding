# go 数据类型

在G0编程语言中，数据类型用于**声明函数和变量**，数据类型的出现是为了把数据分成所需**内存大小**不同的数据，编程的时候需要用大数据的时候才需要申请大内存，就可以**充分利用内存**。 

G0语言按类别有以下几种数据类型：  

1. 布尔型  
    布尔型的值只可以是常量 `true` 或者 `false` 。一个简单的例子：
    `var b bool=true`  
2. 数字类型  
    整型int 和浮点型float32、float64,Go语言支持整型和浮点型数字，并且支持复数，具中位的运算采用补码。
3. 字符串类型  
    字符串就是一串因定长度的字符连接起来的字符序列。G0的字符串是由单个字节连接起来的。Go语言的字符串的字节使用UTF-8编码标识Unicode文本，
4. 派生类型  
    1. 指针类型(Pointer)
    2. 数组类型(c)
    3. 结构体类型(struct)
    4. Channel:类型
    5. 函数类型
    6. 切片类型  
    7. 接口类型  
    8. Map类型  

## 格式化输出类型
```go
package main

import "fmt"

func main() {
	var name string = "tom"
	age := 20
	b := true
	// 格式化输出
	// %T 类型
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", b)
}
```
输出:  
```shell
string
int
bool
```

## 布尔类型

go语言中的布尔类型有两个常量值：`true`和`false`。布尔类型经常用在**条件判断**语句，或者**循环语句**。也可以用在**逻辑表达式**中。

## 数值类型

## 派生类型

### 指针类型
```go
package main

import "fmt"

func main() {
	//指针类型
	a := 100
	p := &a //取a的内存地址 p是指针类型
	fmt.Printf("%T", p) //*int
}
```
<!--stackedit_data:
eyJoaXN0b3J5IjpbLTE1ODY4MzMyMThdfQ==
-->