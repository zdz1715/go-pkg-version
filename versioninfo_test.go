package gopkgversion

import (
	"testing"
)

func TestPrint(t *testing.T) {
	versionInfo := NewVersionInfo("myapp")

	t.Logf("%s\n", versionInfo)
	t.Logf("%s\n", versionInfo.KVString())

	versionInfo = NewVersionInfo()

	t.Logf("%s\n", versionInfo)
	t.Logf("%s\n", versionInfo.KVString())
}
