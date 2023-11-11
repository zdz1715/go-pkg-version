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
		ver := ParseVersion(vi.Version)
		vi.Major = ver.Major()
		vi.Minor = ver.Minor()
	}
	return vi
}

func (vi *VersionInfo) UnsetRuntime() *VersionInfo {
	vi.GoVersion = ""
	vi.Compiler = ""
	vi.Platform = ""
	return vi
}

func (vi *VersionInfo) SetName(name string) *VersionInfo {
	vi.Name = name
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
	latest bool
}

func (v *Version) String() string {
	if v.latest {
		return "latest"
	}

	if v.major == "" {
		return ""
	}
	verStr := "v" + v.major
	if v.minor != "" {
		verStr += "." + v.minor
	}
	return verStr
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

	return v.minor < other.minor
}

func (v *Version) Major() string {
	return v.major
}

func (v *Version) Minor() string {
	return v.minor
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
	majorStart := -1
	majorEnd := 0
	minorStart := 0
	minorEnd := 0
	majorFinish := false
	for i, r := range v {
		if majorStart >= 0 && majorEnd > 0 && minorStart > 0 && minorEnd > 0 {
			break
		}
		if unicode.IsDigit(r) {
			if majorStart < 0 {
				majorStart = i
			}

			if majorFinish && minorStart == 0 {
				minorStart = i
			}

		} else {
			if majorStart >= 0 && majorEnd == 0 {
				majorEnd = i
			}

			if minorStart > 0 && minorEnd == 0 {
				minorEnd = i
			}

			if r == '.' {
				// 重置
				if majorEnd == 0 {
					majorStart = -1
				}

				if minorEnd == 0 {
					minorStart = 0
				}
				if majorEnd > 0 {
					majorFinish = true
				}
			}
		}
	}
	if majorStart < 0 {
		return ver
	}

	if majorEnd < majorStart {
		ver.major = v[majorStart:]
		return ver
	}
	if minorEnd < minorStart {
		ver.major = v[majorStart:majorEnd]
		ver.minor = v[minorStart:]
		return ver
	}

	ver.major = v[majorStart:majorEnd]
	ver.minor = v[minorStart:minorEnd]
	return ver
}
