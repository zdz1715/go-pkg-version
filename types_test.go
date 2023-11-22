package gopkgversion

import (
	"testing"
)

func TestSetVersion(t *testing.T) {
	// 指定版本
	SetVersion("v10.0.1")

	_ = JsonPrint(NewVersionInfo(), nil)
	_ = NamedJsonPrint(NewVersionInfo(), nil)
	_ = PlainTextPrint(NewVersionInfo(), nil)
}
