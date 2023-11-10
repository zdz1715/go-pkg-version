package go_pkg_version

import (
	"fmt"
	"runtime"
)

// Info contains versioning information.
type Info struct {
	Version      string `json:"version"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

// String returns a Go-syntax representation of the Info.
//func (info Info) String() string {
//	jsonString, _ := json.Marshal(info)
//	return string(jsonString)
//}

// Get returns the overall codebase version. It's for detecting
// what code a binary was built from.
func Get() *Info {
	return &Info{
		Version:      version,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
