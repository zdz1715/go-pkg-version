package gopkgversion_test

import (
	gopkgversion "github.com/zdz1715/go-pkg-version"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	t.Logf("%#v", gopkgversion.Get())
	gopkgversion.SetVersion("v1.22.3")
	t.Logf("%#v", gopkgversion.Get())
	gopkgversion.SetVersion("v1.23.2", time.Now())
	gopkgversion.SetName("Kubernetes")
	t.Logf("%s", gopkgversion.Get().Json())
}
