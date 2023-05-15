1. GOROOT  go 的安装目录
2. GOPATH  go 的项目目录， 需要包含bin,srv,pkg
3. 学习内容
   1. 基础语法
   2. 流程控制
   3. 函数
   4. 数组
   5. 指针，结构体
   6. 面向对象
   7. 异常处理
   8. 常用库
   9. 高级
      1.  集合框架，io流，并发编程，网络编程，GUI,反射，编译原理，内存管理，元编程详解

## 配置vsc 支持输入的debug,
如果不配置，输入会不生效 报错 `Unable to process evaluate: debuggee is running` 
详细配置： https://github.com/golang/vscode-go/blob/master/docs/debugging.md#configuration

go 语言学习的官方指南
以前叫这个名字 https://golang.org 现在改名为 https://go.dev/

参考学习文档：

https://golang.iswbm.com/c04/c04_04.html

https://github.com/jiujuan/go-collection

### go语言圣经
https://books.studygolang.com/gopl-zh/

一个go 的web 项目

https://github.com/go-admin-team/go-admin

https://github.com/flipped-aurora/gin-vue-admin

### go 写测试代码
1. 名称需要以_test.go 结尾
2. 入口方法名Test*(t *testing.T) 方式命名

变量可以使用下划线来标识即将被删掉，或者不用的变量
```go
_, ok = m[key]            // map返回2个值

```

1. 基础语法
2. 创建对象，初始化对象，创建接口
3. 项目管理
4. 文件流处理
5. 并发编程
6. 集合操作接口