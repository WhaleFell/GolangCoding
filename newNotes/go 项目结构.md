
# Go 项目结构和go mod最佳实践

## 1. 项目目录
```other
|-- go-test
  |--bin 存放编译后的可执行文件
  |--pkg 存放编译后的包文件
  |--src 存放项目源文件
```
**一般地bin和pkg目录可以不创建，go命令会自动创建（爽否？），只需要创建src目录放代码即可。**

## 2. 环境变量
- GOROOT：安装的`go`路径
- GOPATA：项目的根目录`go-test`   
细心的人注意到，这里有一个`Project GOPATH`，还有一个`Global GOPATH`，把你的项目配置在`Project GOPATH`里，每个项目都不一样，创建另一个项目时这个路径要配置成新项目的。
`Global GOPATH`可以弄一个公共项目，以后就把第三方的包直接装到这里，就可以**自动**在你的项目里引用了。
![输入图片说明](/imgs/2022-11-22/6SavhjegXmvXBois.png)

## .引用包
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
<!--stackedit_data:
eyJoaXN0b3J5IjpbLTE1MTcxMTIzMTUsNjkyNDkwODc4XX0=
-->