package gopkgversion

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	t.Logf("%#v", Get())
	SetVersion("v1.22.3")
	t.Logf("%#v", Get())
	SetVersion("v1.23.2", time.Now())
	t.Logf("%s", Get().Json())
}
