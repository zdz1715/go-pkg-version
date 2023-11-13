package gopkgversion

import (
	"fmt"
	"runtime"
	"unicode"
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
		ver := ParseVersion(vi.Version)
		vi.Major = ver.Major()
		vi.Minor = ver.Minor()
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

type Version struct {
	major  string
	minor  string
	patch  string
	latest bool
}

func (v *Version) String() string {
	if v.latest {
		return "latest"
	}

	if v.major == "" {
		return ""
	}

	if v.minor == "" {
		return v.major
	}

	if v.patch == "" {
		return v.major + "." + v.minor
	}

	return v.major + "." + v.minor + "." + v.patch
}

// Older returns true if this version v is older than the other.
func (v *Version) Older(other *Version) bool {
	if v.latest || other == nil { // Latest is always consider newer, even than future versions.
		return false
	}
	if other.latest {
		return true
	}
	if v.major != other.major {
		return v.major < other.major
	}

	if v.minor != other.minor {
		return v.minor < other.minor
	}

	return v.patch < other.patch
}

func (v *Version) Major() string {
	return v.major
}

func (v *Version) Minor() string {
	return v.minor
}

func (v *Version) Patch() string {
	return v.patch
}

func (v *Version) Latest() bool {
	return v.latest
}

func ParseVersion(v string) *Version {
	ver := new(Version)
	if v == "" {
		return ver
	}
	if v == "latest" {
		ver.latest = true
		return ver
	}
	index := -1
	endIndex := len(v) - 1
	skip := false
	maxLen := 3
	list := make([]string, 0, maxLen)
	for i, r := range v {
		if len(list) >= maxLen {
			break
		}

		newNum := false
		if unicode.IsDigit(r) {
			if index < 0 && !skip {
				index = i
			}
			if i == endIndex && index >= 0 {
				list = append(list, v[index:])
				break
			}
		} else {
			newNum = true
			if r == '.' {
				skip = false
			}
		}

		if newNum && index >= 0 {
			list = append(list, v[index:i])
			index = -1
			if r != '.' {
				skip = true
			}
		}

	}
	if len(list) == 0 {
		return ver
	}
	if len(list) > 0 {
		ver.major = list[0]
	}
	if len(list) > 1 {
		ver.minor = list[1]
	}
	if len(list) > 2 {
		ver.patch = list[2]
	}
	return ver
}
