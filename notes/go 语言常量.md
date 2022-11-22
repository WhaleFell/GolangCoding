# go 语言常量

常量，就是在程序**编译阶段**就确定下来的值，而程序在**运行时则无法改变该值**。在G0程序中，常量可以是数值类型(包括整型、浮点型和复数类型)、布尔类型、字符串类型等。

## 定义常量的语法
定义一个常量使用const关键字，语法格式如下：  
```go
const constantName [type] = value
/* 
const:定义常量关键字
constantName:常量名称
type:常量类型
value:常量的值 
*/
```
### 示例
```go
package main

func main() {
	const PI float64 = 3.14
	const PI2 = 3.1415 //可以省略类型
    // PI2 = 3.15 不能重新对他赋值
	// 批量声明
	const (
		width  = 100
		height = 200
	)

	//多重赋值
	const i, j = 1, 2
	const a, b, c = 1, 2, "foo"
}
```

## `iota` 关键字

`iota`比较特殊，可以被认为是一个可被编译器修改的常量，它默认开始值是`0`，每调用一次加`1`。遇到`const`关键字时被重置为`0`。e.g.  

```go
package main

import "fmt"

func main() {
	// iota 关键字
	const (
		a1 = iota //0
		a2 = iota //类似 i++
		a3 = iota
	)

	fmt.Printf("a1: %v\n", a1) //0
	fmt.Printf("a2: %v\n", a2) //1
	fmt.Printf("a3: %v\n", a3) //2
}
```

### 使用 `_` 跳过某些值  

```go
package main

import "fmt"

func main() {
	// iota 关键字
    // 使用 `_` 跳过某些值
	const (
		a1 = iota //0
		_         //1
		a3 = iota //2
	)

	fmt.Printf("a1: %v\n", a1) //0
	// fmt.Printf("a2: %v\n", a2) //1
	fmt.Printf("a3: %v\n", a3) //2
}
```

### `iota` 声明中间插队

```go
package main

import "fmt"

func main() {
	// iota 关键字
    // `iota` 声明中间插队
	const (
		a1 = iota //0
		a2 = 100         //100
		a3 = iota //1
	)

	fmt.Printf("a1: %v\n", a1) //0
	fmt.Printf("a2: %v\n", a2) //100
	fmt.Printf("a3: %v\n", a3) //1
}
```
<!--stackedit_data:
eyJoaXN0b3J5IjpbMTgwMTE1MDIzMF19
-->