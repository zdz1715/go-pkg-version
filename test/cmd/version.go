package main

import (
	"fmt"
	gopkgversion "github.com/zdz1715/go-pkg-version"
	"os"
)

func main() {
	// 从git tag获取
	gopkgversion.SetVersion("v1.21.1")
	cmd := gopkgversion.NewVersionCommand(&gopkgversion.CmdOptions{
		ParentName: "myapp",
	})
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
