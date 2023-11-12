package gopkgversion

import (
	"fmt"
	"os"
	"testing"
)

func TestNewVersionCommand(t *testing.T) {
	cmd := NewVersionCommand(nil)
	fmt.Println("+ version help")
	_ = cmd.Help()
	fmt.Println("+ version")
	_ = cmd.Execute()
	_ = cmd.Flags().Set("number", "true")
	fmt.Println("+ version --number")
	_ = cmd.Execute()
}

func TestNewVersionCommandWithCmdOptions(t *testing.T) {
	cmd := NewVersionCommand(&CmdOptions{
		Name: "kubectl",
		//NoRuntimeInfo: true,
		// 自定义打印方式
		PrintHandler: func(versionInfo *VersionInfo, opts *PrintOptions) error {
			info := fmt.Sprintf("%s: %s", versionInfo.Name, versionInfo.Version)
			if opts.OnlyNumber {
				_, err := fmt.Println(info)
				return err
			}

			info2 := fmt.Sprintf("go_version: %s platform: %s", versionInfo.GoVersion, versionInfo.Platform)

			_, err := fmt.Fprintln(os.Stdout, info+" "+info2)
			return err
		},
	})
	fmt.Println("+ kubectl version help")
	_ = cmd.Help()
	fmt.Println("+ kubectl version")
	_ = cmd.Execute()

	_ = cmd.Flags().Set("number", "true")
	fmt.Println("+ kubectl version --number")
	_ = cmd.Execute()
}
