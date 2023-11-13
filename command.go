package gopkgversion

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	versionShort   = `Print the version information`
	versionLong    = `Print the version information`
	versionExample = `  # Print the version information
  %s version
  # Print the version number only
  %s version -n 
  %s version --number`
)

type CmdOptions struct {
	Name          string
	NoRuntimeInfo bool
	PrintHandler  PrintHandler
}

// NewVersionCommand prints out the release version info for this command binary.
// It is used as a subcommand of a parent command.

func NewVersionCommand(opts ...*CmdOptions) (*cobra.Command, *VersionInfo) {
	printOpts := new(PrintOptions)
	cmdOpts := new(CmdOptions)
	if len(opts) > 0 && opts[0] != nil {
		cmdOpts = opts[0]
	}
	versionInfo := NewVersionInfo(cmdOpts.Name)
	cmd := &cobra.Command{
		Use:   "version",
		Short: versionShort,
		Long:  versionLong,
		Example: fmt.Sprintf(versionExample,
			cmdOpts.Name,
			cmdOpts.Name,
			cmdOpts.Name,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			vi := versionInfo.Copy()
			if cmdOpts.NoRuntimeInfo {
				vi.UnsetRuntime()
			}
			if cmdOpts.PrintHandler != nil {
				return cmdOpts.PrintHandler(vi, printOpts)
			}
			return NamedJsonPrint(vi, printOpts)
		},
	}

	cmd.Flags().BoolVarP(&printOpts.OnlyNumber, "number", "n", false,
		"Print the version number only")

	return cmd, versionInfo
}
