package gopkgversion

import (
	"fmt"
	"runtime"
)

// VersionInfo contains versioning information.
type VersionInfo struct {
	Name         string `json:"Name,omitempty"`
	Major        string `json:"Major,omitempty"`
	Minor        string `json:"Minor,omitempty"`
	Version      string `json:"Version,omitempty"`
	GitCommit    string `json:"GitCommit,omitempty"`
	GitTreeState string `json:"GitTreeState,omitempty"`
	BuildDate    string `json:"BuildDate,omitempty"`
	GoVersion    string `json:"GoVersion,omitempty"`
	Compiler     string `json:"Compiler,omitempty"`
	Platform     string `json:"Platform,omitempty"`
}

// NewVersionInfo returns the overall codebase version. It's for detecting
// what code a binary was built from.
func NewVersionInfo() *VersionInfo {
	vi := &VersionInfo{
		Version:      version,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,

		GoVersion: runtime.Version(),
		Compiler:  runtime.Compiler,
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
	if vi.Version != "" {
		//versions := strings.Split(vi.Version, ".")
	}
	return vi
}

func (vi *VersionInfo) UnsetRuntime() {
	vi.GoVersion = ""
	vi.Compiler = ""
	vi.Platform = ""
}
