# go-pkg-version
快速设置应用版本信息，如果你需要编写多个应用的版本设置代码，你可能需要这个。

## 安装

```shell
go get -u github.com/zdz1715/go-pkg-version@latest
```

## 使用方式
### 直接设置版本
只需要版本信息，忽略`git`和`buildDate`信息，则可以直接在代码里调用`SetVersion()`方法设置

```go
package main

import (
	"fmt"
	gopkgversion "github.com/zdz1715/go-pkg-version"
)

func main() {
	gopkgversion.SetVersion("v1.22.3")
	info := gopkgversion.NewVersionInfo("golang")
	fmt.Println(info)
}
```
Output:
```shell
golang: {"major":"1","minor":"22","patch":"3","version":"v1.22.3","goVersion":"go1.22.3","compiler":"gc","platform":"windows/amd64"}
```
### 打包注入版本
需要`git`和`buildDate`信息，可以在`Makefile`里注入版本信息，下面使用git tag为版本号
```Makefile
# Git information
GIT_COMMIT = $(shell git rev-parse HEAD)
#GIT_COMMIT_HASH    = $(shell git rev-parse --short HEAD)
GIT_COMMIT_HASH    = $(shell git rev-parse HEAD)
GIT_TAG    = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
GIT_TREESTATE  = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")
BUILDDATE = $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

LDFLAGS += -X github.com/zdz1715/go-pkg-version.version=$(GIT_TAG)
LDFLAGS += -X github.com/zdz1715/go-pkg-version.gitCommit=$(GIT_COMMIT_HASH)
LDFLAGS += -X github.com/zdz1715/go-pkg-version.gitTreeState=$(GIT_TREESTATE)
LDFLAGS += -X github.com/zdz1715/go-pkg-version.buildDate=$(BUILDDATE)

.PHONY: build
build: ## Build binary.
	go build -ldflags "$(LDFLAGS)" -o app app/main.go
```
 
## 版本信息字段

| 字段                 | 说明                           | 
|:-------------------|:-----------------------------|
| name               | 应用名称                         |
| major              | 主要版本号                        |
| minor              | 次要版本号                        |
| version            | 版本号                          |
| gitCommit          | Git 提交hash                   |
| gitTreeState       | Git 提交状态: 'clean' or 'dirty' |
| buildDate          | 构建时间                         |
| goVersion          | go 版本                        |
| compiler           | 编译器名称                        |
| platform           | 系统架构，format: os/arch         |
