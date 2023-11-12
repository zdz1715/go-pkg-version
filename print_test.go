package gopkgversion

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	versionInfo := NewVersionInfo("myapp")

	fmt.Println("> NamedJsonPrint: named json format(default)")
	NamedJsonPrint(versionInfo, nil)
	fmt.Println("> JsonPrint: json format")
	JsonPrint(versionInfo, nil)

	fmt.Println("> PlainTextPrint: plain text format")
	PlainTextPrint(versionInfo, nil)
}
