package gopkgversion

import (
	"reflect"
	"testing"
)

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
			version:       "v0.30.1",
			targetVersion: "v0.1.1",
			want:          false,
		},
		{
			version:       "1.2.1",
			targetVersion: "1.12345.2",
			want:          true,
		},
	}

	for _, tt := range tests {
		flag := Older(tt.version, tt.targetVersion)
		if flag != tt.want {
			t.Errorf("%s < %s = %t, want=%t", tt.version, tt.targetVersion, flag, tt.want)
		}
	}
}

func TestNewVersion(t *testing.T) {
	tests := []struct {
		version string
		want    *Version
	}{
		// [任意字符]主版本号.次版本号.修正版本号[任意字符]
		{version: "1.24.3-20200604", want: New(1, 24, 3)},
		{version: "v1.24.3", want: New(1, 24, 3)},
		{version: "version: 1.24.3", want: New(1, 24, 3)},
		{version: "版本: 1.24.3", want: New(1, 24, 3)},
		{version: "ruby 3.2.2 (2023-03-30 revision e51014f9c0)", want: New(3, 2, 2)},
		{version: "go version go1.21.4 darwin/arm6", want: New(1, 21, 4)},
		// 下面是随意输入的版本号，只查找以'.'分割的每个字符串最先匹配到的数字
		// [任意字符]主版本号[任意字符].次版本号[任意字符][.修正版本号]
		{version: "v52-kfc-4.v50-rc", want: New(52, 50)},
		{version: "v10-2.43-50.30-20220513_rc", want: New(10, 43, 30)},
		{version: "v1.2-20220513_rc", want: New(1, 2)},
		// [任意字符]主版本号[任意字符]
		{version: "v10.", want: New(10)},
		{version: "v10.v", want: New(10)},

		// [任意字符]主版本号[任意字符].次版本号
		{version: "version.100.version.num.23", want: New(100, 23)},
		{version: "1", want: New(1)},
		// 字符无数字
		{version: "latest", want: NewLatest()},
	}

	for _, tt := range tests {
		ver := ParseVersion(tt.version)

		if !reflect.DeepEqual(ver, tt.want) {
			t.Errorf("version=%s target=%s want=%s", tt.version, ver, tt.want)
		}
	}
}

func BenchmarkParseVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseVersion("v1.20.3-2009-03-21")
	}
}
