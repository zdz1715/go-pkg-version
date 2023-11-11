package gopkgversion

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	SetVersion("v1.0.2")
	versionInfo := NewVersionInfo().SetName("kubectl")

	fmt.Println("> NamedJsonPrint: named json format(default)")
	_ = NamedJsonPrint(versionInfo, nil)
	fmt.Println("> JsonPrint: json format")
	_ = JsonPrint(versionInfo, nil)

	fmt.Println("> PlainTextPrint: plain text format")
	_ = PlainTextPrint(versionInfo, nil)
}
