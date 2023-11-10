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

type Options struct {
	ParentName    string
	PrintHandler  PrintHandler
	NoRuntimeInfo bool
}

// NewVersionCommand prints out the release version info for this command binary.
// It is used as a subcommand of a parent command.
func NewVersionCommand(opts *Options) *cobra.Command {
	printOpts := new(PrintOptions)
	if opts == nil {
		opts = new(Options)
	}
	cmd := &cobra.Command{
		Use:   "version",
		Short: versionShort,
		Long:  versionLong,
		Example: fmt.Sprintf(versionExample,
			opts.ParentName,
			opts.ParentName,
			opts.ParentName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			versionInfo := NewVersionInfo()
			versionInfo.Name = opts.ParentName
			if opts.NoRuntimeInfo {
				versionInfo.UnsetRuntime()
			}
			if opts.PrintHandler != nil {
				return opts.PrintHandler(versionInfo, printOpts)
			}
			return JsonPrint(versionInfo, printOpts)
		},
	}

	printOpts.addFlags(cmd.Flags())

	return cmd
}
