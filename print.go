package gopkgversion

import (
	"encoding/json"
	"fmt"
<<<<<<< HEAD
	"github.com/spf13/pflag"
=======
>>>>>>> d18076a (test: git tree state)
	"os"
	"reflect"
	"strings"
)

type PrintOptions struct {
	OnlyNumber bool
}

<<<<<<< HEAD
func (p *PrintOptions) addFlags(flags *pflag.FlagSet) {
	flags.BoolVarP(&p.OnlyNumber, "number", "n", false, "Print the version number only")
}

type PrintHandler func(versionInfo *VersionInfo, opts *PrintOptions) error

var JsonPrint PrintHandler = func(versionInfo *VersionInfo, opts *PrintOptions) error {
	if versionInfo == nil {
		return nil
=======
type PrintHandler func(versionInfo *VersionInfo, opts *PrintOptions) error

func jsonPrint(versionInfo *VersionInfo, opts *PrintOptions, named bool) error {
	if versionInfo == nil {
		return fmt.Errorf("version: nil version information")
>>>>>>> d18076a (test: git tree state)
	}
	if opts != nil && opts.OnlyNumber {
		_, err := fmt.Println(versionInfo.Version)
		return err
	}
<<<<<<< HEAD
	vBytes, err := json.Marshal(versionInfo)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(os.Stdout, string(vBytes))
	return err
}

var NamedJsonPrint PrintHandler = func(versionInfo *VersionInfo, opts *PrintOptions) error {
	if versionInfo == nil {
		return nil
	}
	prefix := ""
	if versionInfo.Name != "" {
		prefix = versionInfo.Name + ": "
		versionInfo.Name = ""
	}
	if opts != nil && opts.OnlyNumber {
		_, err := fmt.Println(prefix + versionInfo.Version)
		return err
=======
	prefix := ""
	if named && versionInfo.Name != "" {
		prefix = versionInfo.Name + ": "
		versionInfo = versionInfo.Copy()
		versionInfo.SetName("")
>>>>>>> d18076a (test: git tree state)
	}

	vBytes, err := json.Marshal(versionInfo)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(os.Stdout, prefix+string(vBytes))
	return err
}

<<<<<<< HEAD
var PlainTextPrint PrintHandler = func(versionInfo *VersionInfo, opts *PrintOptions) error {
	if versionInfo == nil {
		return nil
=======
var JsonPrint PrintHandler = func(versionInfo *VersionInfo, opts *PrintOptions) error {
	return jsonPrint(versionInfo, opts, false)
}

var NamedJsonPrint PrintHandler = func(versionInfo *VersionInfo, opts *PrintOptions) error {
	return jsonPrint(versionInfo, opts, true)
}

var PlainTextPrint PrintHandler = func(versionInfo *VersionInfo, opts *PrintOptions) error {
	if versionInfo == nil {
		return fmt.Errorf("version: nil version information")
>>>>>>> d18076a (test: git tree state)
	}

	if opts != nil && opts.OnlyNumber {
		_, err := fmt.Println(versionInfo.Version)
		return err
	}

	var builder strings.Builder
<<<<<<< HEAD
	builder.WriteString("Version=")
	builder.WriteString(versionInfo.Version)
	vType := reflect.TypeOf(versionInfo).Elem()
	v := reflect.ValueOf(versionInfo).Elem()
	for i := 0; i < vType.NumField(); i++ {
		fieldName := vType.Field(i).Name
		if fieldName == "Version" {
			continue
		}
		value := v.Field(i).String()
		if value != "" {
			builder.WriteString(" ")
			builder.WriteString(fieldName)
			builder.WriteString("=")
			builder.WriteString(value)
=======
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
>>>>>>> d18076a (test: git tree state)
		}

	}
	_, err := fmt.Println(builder.String())
	return err
}
