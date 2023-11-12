package main

import gopkgversion "github.com/zdz1715/go-pkg-version"

func main() {
	// 需要先执行git commit，然后 git tag v0.1.1
	gopkgversion.SetGitInfo(true)
	gopkgversion.SetBuildDate()

	// 在项目任意位置打印版本信息
	gopkgversion.NamedJsonPrint(gopkgversion.NewVersionInfo("myapp"), nil)
}
