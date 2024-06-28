# 环境变量：GOPATH
GOPATH环境变量可以被手动地配置多个路径。 以后，当GOPATH文件
夹被提及的时候，它表示GOPATH环境变量中的第一个路径对应的文件夹。

- GOPATH文件夹中的pkg子文件夹用来缓存被本地项目所依赖的Go模块（一个Go模块为若干
Go库包的集合）的版本。
- GOBIN环境变量用来指定go install子命令产生的Go应用程序二进制可执行文件应该存储
在何处。 它的默认值为GOPATH文件夹中的bin子目录所对应的目录路径。 GOBIN路径需
配置在PATH环境变量中，以便从任意目录运行这些Go应用程序。

# 最简单的Go程序
[实例代码](./first_go_prog.go)
```go
// simplest-go-program.go
package main

func main() {
}

```
## 运行
```
go run simplest-go-program.go
```
如果一个程序的main包中有若干Go源代码文件，我们也可以使用下面的命令运行此程序
```
go run .
```

- `go run` 子命令并不推荐在正式的大项目中使用。`go run` 子命令只是一种方便的方式来运
行简单的Go程序。 对于正式的项目，最好使用 `go build` 或者 `go install` 子命令构建可执
行程序文件来运行Go程序。
- 支持Go模块特性的Go项目的根目录下需要一个 `go.mod` 文件。此文件可以使用 `go mod init`
子命令来生成（见下）。
- 名称以 `_` 和 `.`开头的源代码文件将被Go官方工具链工具忽略掉。

## 更多子命令
- go vet 可以用来检查可能的代码逻辑错误（警告）
- go fmt 格式化Go代码
- go test 进行待援测试和基准测试用例
- go doc 查看Go代码库包的文档

### Go模块特性来简化依赖管理， 对一个支持Go模块特性的项目：
```
go mod init example.com/myproject
```
在当前目录中生成一个go.mod文
件。 当前目录将被视为一个名为example.com/myproject的模块（即当前项目）的根目
录。 此go.mod文件将被用来记录当前项目需要的依赖模块和版本信息。 我们可以手动编
辑或者使用go子命令来修改此文件。
```
go mod tidy
```
通过扫描当前项目中的所有代码来添加未被记录的依赖至go.mod文
件或从go.mod文件中删除不再被使用的依赖。
```
go get
```
拉添加、升级、降级或者删除单个依赖。此命令不如go mod tidy命令常
用。