package gopkgversion

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

// VersionInfo contains versioning information.
type VersionInfo struct {
	name string `json:"-"`

	Major        string `json:"major,omitempty"`
	Minor        string `json:"minor,omitempty"`
	Patch        string `json:"patch,omitempty"`
	Version      string `json:"version,omitempty"`
	GitCommit    string `json:"gitCommit,omitempty"`
	GitTreeState string `json:"gitTreeState,omitempty"`
	BuildDate    string `json:"buildDate,omitempty"`

	GoVersion string `json:"goVersion,omitempty"`
	Compiler  string `json:"compiler,omitempty"`
	Platform  string `json:"platform,omitempty"`
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
		vi.Major = strconv.Itoa(ver.Major())
		vi.Minor = strconv.Itoa(ver.Minor())
		vi.Patch = strconv.Itoa(ver.Patch())
	}

	if len(name) > 0 {
		vi.name = name[0]
	}

	return vi
}

func (vi *VersionInfo) String() string {
	var builder strings.Builder
	if vi.name != "" {
		builder.WriteString(vi.name + ": ")
	}
	vBytes, err := json.Marshal(vi)
	if err == nil {
		builder.Write(vBytes)
	}
	return builder.String()
}

func (vi *VersionInfo) GetName() string {
	return vi.name
}

func (vi *VersionInfo) KVString() string {
	var builder strings.Builder
	vType := reflect.TypeOf(vi).Elem()
	v := reflect.ValueOf(vi).Elem()
	length := vType.NumField()
	for i := 0; i < length; i++ {
		fieldName := vType.Field(i).Name
		value := v.Field(i).String()
		if value != "" {
			builder.WriteString(strings.ToLower(fieldName[:1]) + fieldName[1:])
			builder.WriteString("=")
			builder.WriteString(value)
			if i < length-1 {
				builder.WriteString(" ")
			}
		}

	}
	return builder.String()
}

func (vi *VersionInfo) UnsetRuntime() *VersionInfo {
	vi.GoVersion = ""
	vi.Compiler = ""
	vi.Platform = ""
	return vi
}

func (vi *VersionInfo) Copy(names ...string) *VersionInfo {
	name := vi.name
	if len(names) > 0 {
		name = names[0]
	}
	return &VersionInfo{
		name:         name,
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
