package gopkgversion

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

// Base version information.
//
// This is the fallback data used when version information from git is not
// provided via go ldflags. It provides an approximation of the version for
// ad-hoc builds (e.g. `go build`) that cannot get the version
// information from git.
var (
	// if no version is set, get it from git tag,
	// output of git describe --tags --abbrev=0 --exact-match
	version = "v0.0.0"

	// sha1 from git, output of $(git rev-parse HEAD)
	gitCommit = "" // sha1 from git, output of $(git rev-parse HEAD)

	// state of git tree, either "clean" or "dirty"
	// output of $(test -n "`git status --porcelain`" && echo "dirty" || echo "clean")
	gitTreeState = "" // state of git tree, either "clean" or "dirty"

	// build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
	buildDate = ""
)

func SetVersion(v ...string) {
	setTagToVersion := true
	if len(v) > 0 && v[0] != "" {
		version = v[0]
		setTagToVersion = false
	}
	setGitInfo(setTagToVersion)
	setBuildDate()
}

func setGitInfo(tagToVersion bool) {
	git, err := exec.LookPath("git")
	if err != nil {
		fmt.Printf("version: %s", err)
		return
	}
	hashCmd := exec.Command(git, "rev-parse", "HEAD")
	if out, err := hashCmd.Output(); err == nil {
		gitCommit = string(bytes.TrimSpace(out))
	}
	stateCmd := exec.Command("/bin/sh", "-c", "test -n \"`git status --porcelain`\" && echo dirty || echo clean")
	if out, err := stateCmd.Output(); err == nil {
		gitTreeState = string(bytes.TrimSpace(out))
	}
	if tagToVersion {
		tagCmd := exec.Command(git, "describe", "--tags", "--abbrev=0", "--exact-match")
		if out, err := tagCmd.Output(); err == nil {
			version = string(bytes.TrimSpace(out))
		}
	}
}

func setBuildDate() {
	buildDate = time.Now().In(time.UTC).Format(time.RFC3339)
}
