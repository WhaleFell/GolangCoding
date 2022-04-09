# go标识符-关键字-命名规则

## 标识符
标识符的英文是identifier,通俗的讲，就是给变量、常量、函数、方法、结构体、数组、切片、接口起名
字。  

标识符的组成:  
1. 标识符由数字、字母和下划线(_) 组成。`123abc_`  
2. 只能以字母和下划线(_)开头。`abc123_sysVar123abc`    
3. 标识符区分大小写。`nameName NAMET`  

## 标识符的命名

```go
package main

func main() {
	// 正确的标识符
	var name string
	var age int
	var _sys int

	/*
		错误的标识符
		var 1name string
		var &age int
	*/

}
```

## Go 关键字

## go语言命名规范  

Go是一门区分大小写的语言  
命名规则涉及变量、常量、全局函数、结构、接口、方法等的命名。G0语言从语法层面进行了以下限定：任何需要对外暴露的名字必须以大写字母开头，不需要对外暴露的则应该以小写字母开头。  
当命名(包括常量、变量、类型、函数名、结构字段等等)以一个大写字母开头，如：GetUserName,那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出(像面向对象语言中的public)；命名如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的(像面向对象语言中的private)  

### 包名称

保持package的名字和目录保持一致，尽量采取有意义的包名，简短，有意义，尽量和标准库不要冲突。包名应该为小写单词，不要使用下划线或者混合大小写。  

```go
package dao  
package service
```

### 文件命名

尽量采取有意义的文件名，简短，有意义，应该为小写单词，使用下划线分隔各个单词。`customer_dao.go`

### 结构体命名  

采用驼峰命名法，首字母根据访问控制大写或者小写 `struct` 申明和初始化格式采用多行，例如下面：

```go
type CustomerOrder struct{
Name string
Address string
}
order:=CustomerOrder{"tom","北京海淀"}
```

### 接口命名  

命名规则基本和上面的结构体类型
单个函数的结构名以“er”作为后缀，例如Reader,Writer.

```go
type Reader interface{
Read(p [] byte)(n int,err error)
}
```

### 变量命名

和结构体类似，变量名称一般遵循驼峰法，首字母根据访问控制原则大写或者小写，但遇到特有名词时，需要遵循以下规则：  
如果变量为私有，且特有名词为首个单词，则使用小写，如appService若变量类型为bool类型，则名称应以Has,Is,Can或Allow开头  

```go
var isExist bool
var hasConflict bool
var canManage bool
var allowGitHook bool
```