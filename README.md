# go-pkg-version
快速设置应用版本信息，如果你需要编写多个应用的版本设置代码，你可能需要这个。

## 安装

```shell
go get -u github.com/zdz1715/go-pkg-version@latest
```

## 使用方式
- [example/not-use-command](./example/not-use-command/main.go)
- [example/use-command](./example/use-command/main.go)
- [example/tag-to-version](./example/tag-to-version/main.go)
 
## 版本信息字段

| 字段                 | 说明                           | 
|:-------------------|:-----------------------------|
| name               | 应用名称，需要手动设置                  |
| major              | 主要版本号，根据版本号自动获取              |
| minor              | 次要版本号，根据版本号自动获取              |
| version            | 版本号，可自动获取git tag为版本号或者手动设置   |
| gitCommit          | Git 提交hash                   |
| gitTreeState       | Git 提交状态: 'clean' or 'dirty' |
| buildDate          | 构建时间                         |
| goVersion          | go 版本                        |
| compiler           | 编译器名称                        |
| platform           | 系统架构，format: os/arch         |

## 打印格式
### 内置的打印方法

```go
package main

import gopkgversion "github.com/zdz1715/go-pkg-version"

func main() {
	versionInfo := gopkgversion.NewVersionInfo("myapp")
	
	// NamedJsonPrint output:
	// myapp: {"major":"0","minor":"0","version":"v0.0.0","goVersion":"go1.21.4","compiler":"gc","platform":"darwin/arm64"}
	gopkgversion.NamedJsonPrint(versionInfo, nil)

	// JsonPrint output:
	// {"name":"myapp","major":"0","minor":"0","version":"v0.0.0","goVersion":"go1.21.4","compiler":"gc","platform":"darwin/arm64"}
	gopkgversion.JsonPrint(versionInfo, nil)

	// PlainTextPrint output:
	// name=myapp major=0 minor=0 version=v0.0.0 goVersion=go1.21.4 compiler=gc platform=darwin/arm64
	gopkgversion.PlainTextPrint(versionInfo, nil)
}
```
### 自定义打印方法

```go
package main

import (
	"fmt"
	gopkgversion "github.com/zdz1715/go-pkg-version"
)

var CustomePrint gopkgversion.PrintHandler = func(versionInfo *gopkgversion.VersionInfo, opts *gopkgversion.PrintOptions) error {
    fmt.Println(versionInfo.Version)
    return nil
}

func main()  {
	versionInfo := gopkgversion.NewVersionInfo("myapp")

	CustomePrint(versionInfo, nil)
}
```
