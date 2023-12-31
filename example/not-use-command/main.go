package main

import (
	gopkgversion "github.com/zdz1715/go-pkg-version"
)

func main() {
	gopkgversion.SetVersion("v1.10.1")

	// 在项目任意位置打印版本信息
	gopkgversion.NamedJsonPrint(gopkgversion.NewVersionInfo("myapp").UnsetRuntime(), nil)
	gopkgversion.NamedJsonPrint(gopkgversion.NewVersionInfo("myapp"), nil)
}
