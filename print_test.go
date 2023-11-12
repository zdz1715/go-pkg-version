package gopkgversion

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	versionInfo := NewVersionInfo("kubectl")

	fmt.Println("> NamedJsonPrint: named json format(default)")
	_ = NamedJsonPrint(versionInfo, nil)
	fmt.Println("> JsonPrint: json format")
	_ = JsonPrint(versionInfo, nil)

	fmt.Println("> PlainTextPrint: plain text format")
	_ = PlainTextPrint(versionInfo, nil)
}
