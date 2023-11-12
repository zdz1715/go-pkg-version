package gopkgversion

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type PrintOptions struct {
	OnlyNumber bool
}

type PrintHandler func(versionInfo *VersionInfo, opts *PrintOptions) error

func jsonPrint(versionInfo *VersionInfo, opts *PrintOptions, named bool) error {
	if versionInfo == nil {
		return fmt.Errorf("version: nil version information")
	}
	if opts != nil && opts.OnlyNumber {
		_, err := fmt.Println(versionInfo.Version)
		return err
	}
	prefix := ""
	if named && versionInfo.Name != "" {
		prefix = versionInfo.Name + ": "
		versionInfo = versionInfo.Copy()
		versionInfo.Name = ""
	}

	vBytes, err := json.Marshal(versionInfo)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(os.Stdout, prefix+string(vBytes))
	return err
}

var JsonPrint PrintHandler = func(versionInfo *VersionInfo, opts *PrintOptions) error {
	return jsonPrint(versionInfo, opts, false)
}

var NamedJsonPrint PrintHandler = func(versionInfo *VersionInfo, opts *PrintOptions) error {
	return jsonPrint(versionInfo, opts, true)
}

var PlainTextPrint PrintHandler = func(versionInfo *VersionInfo, opts *PrintOptions) error {
	if versionInfo == nil {
		return fmt.Errorf("version: nil version information")
	}

	if opts != nil && opts.OnlyNumber {
		_, err := fmt.Println(versionInfo.Version)
		return err
	}

	var builder strings.Builder

	vType := reflect.TypeOf(versionInfo).Elem()
	v := reflect.ValueOf(versionInfo).Elem()
	length := vType.NumField()
	for i := 0; i < length; i++ {
		fieldName := vType.Field(i).Name
		value := v.Field(i).String()
		if value != "" {
			builder.WriteString(strings.ToLower(fieldName[:1]) + fieldName[1:])
			builder.WriteString("=")
			builder.WriteString(value)
			if i < length-1 {
				builder.WriteString(" ")
			}
		}

	}
	_, err := fmt.Println(builder.String())
	return err
}
