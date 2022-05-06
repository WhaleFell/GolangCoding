# golang 变量

变量是计算机语言中能储存计算结果或能表示值的抽象概念。不同的变量保存的数据类型可能会不一样。  

## 声明变量

G0语言中的**变量需要声明后才能使用**，同一作用域内不支持重复声明。并且G0语言的**变量声明后必须使用**。

## 声明变量的语法
```go
var identifier type
/* 
var:声明变量关键字
1dent1fier:变量名称
type:变量类型
*/
```
e.g.  
```go
package main

import "fmt"

func main() {
	var name string
	fmt.Printf("name: %v\n", name)
}
```

## 批量声明
```go
package main

import "fmt"

func main() {
	// 批量声明
	var (
		name string
		age  int
		b    bool
	)

	name = "hyy"
	age = 14
	b = false

	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("b: %v\n", b)
}
```

## 变量的初始化

G0语言在声明变量的时候，会**自动对变量对应的内存区域进行初始化操作**。每个变量会被初始化成其类型的**默认值**，例如：整型和浮点型变量的默认值为`0`。字符串变量的默认值为空字符串`""`。布尔型变量默认为`false`,切片、函数、指针变量的默认为`nil`.

### 变量初始化语法

```go
var 变量名 类型 = 表达式
```

e.g.  
```go
package main

import "fmt"

func main() {
	//变量初始化 赋初始值
	var title string = "lovehyy"
	var site string = "http://baidu.com"
	var agge int = 30
}
```

### 类型推断

```go
package main

import "fmt"

func main() {
	// 类型推断
	var name = "hyy"
	var age = 20
	var b = true
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("b: %v\n", b)
}
```

### 批量初始化
```go
package main

import "fmt"

func main() {
	//批量初始化
	var name, age, b = "tom", 20, true
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("b: %v\n", b)
}
```

## 短变量声明

在函数内部，可以使用`:=`运算符对变量进行声明和初始化。  

```go
package main

import "fmt"

// name := "hyy" 不能用于函数外部

func main() {
	//短变量声明 := 不能用于函数外部
	name := "hyy"
	age := 51
	b := true
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("b: %v\n", b)
}
```

> **注意：这种方法只适合在函数内部，函数外面不能使用。**

## 匿名变量

如果我们接收到**多个变量**，有一些**变量使用不到**，可以**使用下划线_表示变量名称**，这种变量叫做匿名变量。e.g.  

1. 函数定义  
```go
func getNameAndAge() (name string, age int) {
	return "hyy", 30
}
```  
等价于： 
```go
func getNameAndAge() (string, int) {
	return "hyy", 30
}
```  

2. 接收函数返回值
```go
package main

import "fmt"

func getNameAndAge() (string, int) {
	return "hyy", 30
}
func main() {
	// 调用函数,匿名变量,丢弃age
	name, _ := getNameAndAge()
	fmt.Printf("name: %v\n", name)
}
```