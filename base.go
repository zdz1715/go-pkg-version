package gopkgversion

import (
	"bytes"
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
	// output of $(git describe --tags --abbrev=0 --exact-match)
	version = "v0.0.0"

	// sha1 from git, output of $(git rev-parse HEAD)
	gitCommit = "" // sha1 from git, output of $(git rev-parse HEAD)

	// state of git tree, either "clean" or "dirty"
	// output of $(test -n "`git status --porcelain`" && echo "dirty" || echo "clean")
	gitTreeState = "" // state of git tree, either "clean" or "dirty"

	// build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
	buildDate = ""
)

func SetVersion(ver string) {
	version = ver
}

func SetGitInfo(tagToVersion bool, dir ...string) {
	git, err := exec.LookPath("git")
	if err != nil {
		return
	}
	gitRepoDir := ""
	if len(dir) > 0 && dir[0] != "" {
		gitRepoDir = dir[0]
	}
	gitWorkCmd := exec.Command(git, "worktree", "list")
	gitWorkCmd.Dir = gitRepoDir
	if _, err = gitWorkCmd.Output(); err != nil {
		return
	}
	hashCmd := exec.Command("/bin/sh", "-c", "git rev-parse HEAD")
	hashCmd.Dir = gitRepoDir
	if out, err := hashCmd.Output(); err == nil {
		gitCommit = string(bytes.TrimSpace(out))
	}
	stateCmd := exec.Command("/bin/sh", "-c", "test -n \"`git status --porcelain`\" && echo dirty || echo clean")
	stateCmd.Dir = gitRepoDir
	if out, err := stateCmd.Output(); err == nil {
		gitTreeState = string(bytes.TrimSpace(out))
	}
	if tagToVersion {
		tagCmd := exec.Command("/bin/sh", "-c", "git describe --tags --abbrev=0 --exact-match")
		tagCmd.Dir = gitRepoDir
		if out, err := tagCmd.Output(); err == nil {
			version = string(bytes.TrimSpace(out))
		}
	}
}

func SetBuildDate(ts ...string) {
	if len(ts) > 0 && ts[0] != "" {
		buildDate = ts[0]
		return
	}
	buildDate = time.Now().In(time.UTC).Format(time.RFC3339)
}
