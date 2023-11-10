package gopkgversion

import (
	"fmt"
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

func TestNewVersionCommandByParentName(t *testing.T) {
	cmd := NewVersionCommand(&Options{
		ParentName: "kubectl",
	})
	fmt.Println("+ kubectl version help")
	_ = cmd.Help()
	fmt.Println("+ kubectl version")
	_ = cmd.Execute()

	_ = cmd.Flags().Set("number", "true")
	fmt.Println("+ kubectl version --number")
	_ = cmd.Execute()
}

func TestNewVersionCommandByCustomPrint(t *testing.T) {
	cmd := NewVersionCommand(&Options{
		ParentName:    "kubectl",
		NoRuntimeInfo: true,
		//PrintHandler: PlainTextPrint,
		PrintHandler: NamedJsonPrint,
		//PrintHandler: func(versionInfo *VersionInfo, flags *CmdFlags) error {
		//	info := fmt.Sprintf("%s: %s", versionInfo.Name, versionInfo.Version)
		//	if flags.OnlyNumber {
		//		_, err := fmt.Println(info)
		//		return err
		//	}
		//
		//	info2 := fmt.Sprintf("go_version: %s platform: %s", versionInfo.GoVersion, versionInfo.Platform)
		//
		//	_, err := fmt.Fprintln(os.Stdout, info+" "+info2)
		//	return err
		//},
	})
	fmt.Println("+ kubectl version help")
	_ = cmd.Help()
	fmt.Println("+ kubectl version")
	_ = cmd.Execute()

	_ = cmd.Flags().Set("number", "true")
	fmt.Println("+ kubectl version --number")
	_ = cmd.Execute()
}

func TestSetVersion(t *testing.T) {
	SetVersion("v0.0.1")
	_ = JsonPrint(NewVersionInfo(), nil)
}
