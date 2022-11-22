
# Go 项目结构

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

<!--stackedit_data:
eyJoaXN0b3J5IjpbMjAyMzYxNDUwNSw2OTI0OTA4NzhdfQ==
-->