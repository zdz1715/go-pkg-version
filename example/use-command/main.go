package main

import (
	"fmt"
	"github.com/spf13/cobra"
	gopkgversion "github.com/zdz1715/go-pkg-version"
	"os"
)

var rootCmd = cobra.Command{
	Use: "myapp",
}

func main() {
	versionCmd, _ := gopkgversion.NewVersionCommand(&gopkgversion.CmdOptions{
		Name: "myapp",
	})
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
