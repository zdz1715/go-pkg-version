package gopkgversion

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
)

// Info contains versioning information.
type Info struct {
	Name         string `json:"name,omitempty"`
	Major        string `json:"major,omitempty"`
	Minor        string `json:"minor,omitempty"`
	Patch        string `json:"patch,omitempty"`
	Version      string `json:"version,omitempty"`
	GitCommit    string `json:"gitCommit,omitempty"`
	GitTreeState string `json:"gitTreeState,omitempty"`
	BuildDate    string `json:"buildDate,omitempty"`
	GoVersion    string `json:"goVersion,omitempty"`
	Compiler     string `json:"compiler,omitempty"`
	Platform     string `json:"platform,omitempty"`
}

func Get() Info {
	ver := ParseVersion(version)
	return Info{
		Name:         name,
		Major:        strconv.Itoa(ver.Major()),
		Minor:        strconv.Itoa(ver.Minor()),
		Patch:        strconv.Itoa(ver.Patch()),
		Version:      version,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

func (v Info) String() string {
	return v.Version
}

func (v Info) Json() string {
	bs, _ := json.Marshal(v)
	return string(bs)
}
