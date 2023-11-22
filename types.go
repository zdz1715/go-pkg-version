package gopkgversion

import (
	"fmt"
	"runtime"
	"strconv"

	versionPkg "github.com/zdz1715/go-utils/version"
)

// VersionInfo contains versioning information.
type VersionInfo struct {
	Name         string `json:"name,omitempty"`
	Major        string `json:"major,omitempty"`
	Minor        string `json:"minor,omitempty"`
	Version      string `json:"version,omitempty"`
	GitCommit    string `json:"gitCommit,omitempty"`
	GitTreeState string `json:"gitTreeState,omitempty"`
	BuildDate    string `json:"buildDate,omitempty"`
	GoVersion    string `json:"goVersion,omitempty"`
	Compiler     string `json:"compiler,omitempty"`
	Platform     string `json:"platform,omitempty"`
}

// NewVersionInfo returns the overall codebase version. It's for detecting
// what code a binary was built from.
func NewVersionInfo(name ...string) *VersionInfo {
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
		ver := versionPkg.ParseVersion(vi.Version)
		vi.Major = strconv.Itoa(ver.Major())
		vi.Minor = strconv.Itoa(ver.Minor())
	}
	if len(name) > 0 && name[0] != "" {
		vi.Name = name[0]
	}
	return vi
}

func (vi *VersionInfo) UnsetRuntime() *VersionInfo {
	vi.GoVersion = ""
	vi.Compiler = ""
	vi.Platform = ""
	return vi
}

func (vi *VersionInfo) Copy() *VersionInfo {
	return &VersionInfo{
		Name:         vi.Name,
		Major:        vi.Major,
		Minor:        vi.Minor,
		Version:      vi.Version,
		GitCommit:    vi.GitCommit,
		GitTreeState: vi.GitTreeState,
		BuildDate:    vi.BuildDate,
		GoVersion:    vi.GoVersion,
		Compiler:     vi.Compiler,
		Platform:     vi.Platform,
	}
}
