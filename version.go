package gopkgversion

import (
	"strconv"
	"strings"
	"unicode"
)

type status uint8

const (
	majorSet status = 1 << iota
	minorSet
	patchSet

	allSet = majorSet | minorSet | patchSet
)

type Version struct {
	major int
	minor int
	patch int

	latest bool
	status status
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
	for i, r := range v {
		if ver.complete() {
			break
		}
		newNum := false
		if unicode.IsDigit(r) {
			if index < 0 && !skip {
				index = i
			}
			if i == endIndex && index >= 0 {
				ver.add(v[index:])
				break
			}
		} else {
			newNum = true
			if r == '.' {
				skip = false
			}
		}

		if newNum && index >= 0 {
			ver.add(v[index:i])
			index = -1
			if r != '.' {
				skip = true
			}
		}

	}

	return ver
}

// New a version number, the parameter order is the main version number, the second version number,
// the patch version number
func New(num ...int) *Version {
	ver := new(Version)
	for _, n := range num {
		if ver.complete() {
			break
		}
		ver.addInt(n)
	}
	return ver
}

func NewLatest() *Version {
	return &Version{
		latest: true,
	}
}

func (v *Version) addInt(n int) {
	if v.status&majorSet == 0 {
		v.major = n
		v.status = majorSet
		return
	}

	if v.status&minorSet == 0 {
		v.minor = n
		v.status = majorSet | minorSet
		return
	}

	if v.status&patchSet == 0 {
		v.patch = n
		v.status = allSet
		return
	}
}

func (v *Version) add(str string) {
	n, err := strconv.Atoi(str)
	if err == nil {
		v.addInt(n)
	}
}

func (v *Version) complete() bool {
	return v.status == allSet
}

// String Return the original version number.
// e.g. 0.0.0
func (v *Version) String() string {
	if v.latest {
		return "latest"
	}

	if v.status&majorSet == 0 {
		return ""
	}

	builder := new(strings.Builder)
	builder.WriteString(strconv.Itoa(v.major))
	if v.status&minorSet != 0 {
		builder.WriteByte('.')
		builder.WriteString(strconv.Itoa(v.minor))
	}
	if v.status&patchSet != 0 {
		builder.WriteByte('.')
		builder.WriteString(strconv.Itoa(v.patch))
	}
	return builder.String()
}

// Version Return the version number with prefix.
// e.g. v0.0.0
func (v *Version) Version(prefix ...string) string {
	if v.latest {
		return "latest"
	}
	str := v.String()
	if len(prefix) > 0 {
		return prefix[0] + str
	}
	return "v" + str
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

func (v *Version) Major() int {
	return v.major
}

func (v *Version) Minor() int {
	return v.minor
}

func (v *Version) Patch() int {
	return v.patch
}

func (v *Version) Latest() bool {
	return v.latest
}

func Older(v1, v2 string) bool {
	return ParseVersion(v1).Older(ParseVersion(v2))
}
