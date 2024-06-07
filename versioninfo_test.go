package gopkgversion

import (
	"testing"
)

func TestPrint(t *testing.T) {
	SetVersion("v1.22.3")
	versionInfo := NewVersionInfo("golang")

	t.Logf("%s\n", versionInfo)
	t.Logf("%s\n", versionInfo.KVString())

	versionInfo = NewVersionInfo()

	t.Logf("%s\n", versionInfo)
	t.Logf("%s\n", versionInfo.KVString())
}
