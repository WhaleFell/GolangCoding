
# Go 项目结构和go mod最佳实践

## 项目目录
```other
|-- go-test
  |--bin 存放编译后的可执行文件
  |--pkg 存放编译后的包文件
  |--src 存放项目源文件
```
**一般地bin和pkg目录可以不创建，go命令会自动创建（爽否？），只需要创建src目录放代码即可。**

## 环境变量
- GOROOT：安装的`go`路径
- GOPATA：项目的根目录`go-test`   
细心的人注意到，这里有一个`Project GOPATH`，还有一个`Global GOPATH`，把你的项目配置在`Project GOPATH`里，每个项目都不一样，创建另一个项目时这个路径要配置成新项目的。
`Global GOPATH`可以弄一个公共项目，以后就把第三方的包直接装到这里，就可以**自动**在你的项目里引用了。
![输入图片说明](/imgs/2022-11-22/6SavhjegXmvXBois.png)

## 引用包
```tree
|____src
| |____main
| | |____calc
| | | |____add.go
| | |____main.go
```
![输入图片说明](/imgs/2022-11-22/hrvDgSCo9FodCkKa.png)
### 注意点：  
1.  `add.go`中的Add函数名首字母必须大写，只有大写的才是Public权限，外面的包才能访问，否则只能自己文件夹下代码才能访问。
2.  `add.go`的改名为addyyy.go也可以，查找add包的时候，**并不会根据add.go这个文件名来查找**。而是根据文件夹名来查找，一个文件夹下的所有文件都属于同一个包。所以函数变量自然不能重复。
3.  `main`中调用add.Add(1,2)时，add是包， 必须跟`add.go`中的`package`处的包名一致，否则报错。
4.  import后， 根据`GOROOT`和`GOPATH`查找对应的包，**src这个目录名可不是能随便取的**。
### 引用第三方项目
![输入图片说明](/imgs/2022-11-22/oZ1JJ9bmA0fF9w5h.png)

## Go mod 包管理工具

### Why?

> 1. 使用go mod仓库中可以不用再上传依赖代码包，防止代码仓库过大浪费以及多个项目同时用包时的浪费
> 2. 可以管理引用包的版本，这一点是gopath（src模式）和`vendor`做不到的
> 3. 如果依赖 `gopath` 不同项目如果引用了同一个软件包的不同版本，就会造成编译麻烦

gopath是go之前的默认策略，每个项目在运行时都要严格放在`src`目录下，而go mod不用。

`Go mod`的优点：
-   自动下载依赖包
-   项目不必放在`$GOPATH/src`内了
-   项目内会生成一个`go.mod`文件，列出包依赖
-   第三方包会准确的指定版本号
-   对于已经转移的包，可以用 `replace` 申明替换，不需要改代码

### 配置
```shell
# 设置环境变量
go env -w GO111MODULE="on"
# china proxy
go env -w GOPROXY=https://goproxy.io
# 初始化 mod 项目
go mod init [packetName]
# add missing and remove unused modules
# 自动增加包和删除无用包到 GOPATH 目录下
go mod tidy
```
注意：只要在本地**设置一个公用path目录**就可以了，**全部的包都会下载到那里**，其他本地项目用到时就可以共享了

自动生成了 `go.mod` 和 `go.sum` 文件，可以不用理会，下面是简单介绍.

### go.mod
```go
// 模块的引用路径
module github.com/BingmingWong/module-test 

// 项目使用的 go 版本
go 1.14

// 项目所需的直接依赖包及其版本
require (
example.com/apple v0.1.2
)

// 忽略指定版本的依赖包
exclude example.com/banana v1.2.4

// 由于在国内访问golang.org/x的各个包都需要翻墙，你可以在go.mod中使用replace替换成github上对应的库。
replace（
golang.org/x/text v0.3.0 = > github.com/golang/text v0.3.0
)

```

<!--stackedit_data:
eyJoaXN0b3J5IjpbLTEwOTQyOTgzMDYsNjkyNDkwODc4XX0=
-->