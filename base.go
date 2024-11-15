package gopkgversion

import "time"

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

func SetVersion(ver string, t ...time.Time) {
	version = ver
	if len(t) > 0 {
		buildDate = t[0].Format(time.RFC3339)
	}
}
