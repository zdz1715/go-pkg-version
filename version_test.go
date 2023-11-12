package gopkgversion

import (
	"reflect"
	"testing"
)

func BenchmarkParseVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseVersion("v1.20.3-2009-03-21")
	}
}

func TestVersion_Older(t *testing.T) {
	tests := []struct {
		version       string
		targetVersion string
		want          bool
	}{
		{
			version:       "v0.2.1",
			targetVersion: "v0.3.1",
			want:          true,
		},
		{
			version:       "v0.3.1",
			targetVersion: "v0.1.1",
			want:          false,
		},
		// 只比较major.minor
		{
			version:       "1.4.1",
			targetVersion: "1.4.2",
			want:          true,
		},
	}

	for _, tt := range tests {
		ver := ParseVersion(tt.version)
		targetVer := ParseVersion(tt.targetVersion)
		flag := ver.Older(targetVer)
		if flag != tt.want {
			t.Errorf("%s < %s = %t, want=%t", tt.version, tt.targetVersion, flag, tt.want)
		}
	}
}

func TestParseVersion(t *testing.T) {

	tests := []struct {
		version string
		want    *Version
	}{
		// [任意字符]主版本号.次版本号.修正版本号[任意字符]
		{version: "1.24.3", want: &Version{
			major: "1",
			minor: "24",
		}},
		{version: "v1.24.3", want: &Version{
			major: "1",
			minor: "24",
		}},
		{version: "version: 1.24.3", want: &Version{
			major: "1",
			minor: "24",
		}},
		{version: "版本: 1.24.3", want: &Version{
			major: "1",
			minor: "24",
		}},
		{version: "ruby 3.2.2 (2023-03-30 revision e51014f9c0)", want: &Version{
			major: "3",
			minor: "2",
		}},
		{version: "go version go1.21.4 darwin/arm6", want: &Version{
			major: "1",
			minor: "21",
		}},
		// 下面是随意输入的版本号，只查找以'.'分割的每个字符串最先匹配到的数字
		// [任意字符]主版本号[任意字符].次版本号[任意字符][.修正版本号]
		{version: "version.52-kfc-4.v50-rc", want: &Version{
			major: "52",
			minor: "50",
		}},
		{version: "v10-2.43-50.3.20220513_rc", want: &Version{
			major: "10",
			minor: "43",
		}},
		{version: "v1.2-20220513_rc", want: &Version{
			major: "1",
			minor: "2",
		}},
		// [任意字符]主版本号[任意字符]
		{version: "v100.", want: &Version{
			major: "100",
			minor: "",
		}},
		{version: "v10.v", want: &Version{
			major: "10",
			minor: "",
		}},

		// [任意字符]主版本号[任意字符].次版本号
		{version: "version.100.version.num.23", want: &Version{
			major: "100",
			minor: "23",
		}},
		// 字符无数字
		{version: "latest", want: &Version{
			latest: true,
		}},
	}

	for _, tt := range tests {
		ver := ParseVersion(tt.version)

		if !reflect.DeepEqual(ver, tt.want) {
			t.Errorf("version=%s target=%v want=%v", tt.version, ver, tt.want)
		}
	}
}

func TestSetVersion(t *testing.T) {
	// 指定版本
	SetVersion("v0.0.1")
	SetBuildDate()
	_ = JsonPrint(NewVersionInfo(), nil)

	// 从git tag获取，先执行: git tag v1.0.1
	SetGitInfo(true)
	_ = JsonPrint(NewVersionInfo(), nil)
}
