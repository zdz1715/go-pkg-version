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

<<<<<<< HEAD
type Options struct {
=======
type CmdOptions struct {
>>>>>>> d18076a (test: git tree state)
	ParentName    string
	PrintHandler  PrintHandler
	NoRuntimeInfo bool
}

// NewVersionCommand prints out the release version info for this command binary.
// It is used as a subcommand of a parent command.
<<<<<<< HEAD
func NewVersionCommand(opts *Options) *cobra.Command {
	printOpts := new(PrintOptions)
	if opts == nil {
		opts = new(Options)
=======
func NewVersionCommand(opts ...*CmdOptions) *cobra.Command {
	printOpts := new(PrintOptions)
	cmdOpts := new(CmdOptions)
	if len(opts) > 0 && opts[0] != nil {
		cmdOpts = opts[0]
>>>>>>> d18076a (test: git tree state)
	}
	cmd := &cobra.Command{
		Use:   "version",
		Short: versionShort,
		Long:  versionLong,
		Example: fmt.Sprintf(versionExample,
<<<<<<< HEAD
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
=======
			cmdOpts.ParentName,
			cmdOpts.ParentName,
			cmdOpts.ParentName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			versionInfo := NewVersionInfo().SetName(cmdOpts.ParentName)
			if cmdOpts.NoRuntimeInfo {
				versionInfo.UnsetRuntime()
			}
			if cmdOpts.PrintHandler != nil {
				return cmdOpts.PrintHandler(versionInfo, printOpts)
			}
			return NamedJsonPrint(versionInfo, printOpts)
		},
	}

	cmd.Flags().BoolVarP(&printOpts.OnlyNumber, "number", "n", false,
		"Print the version number only")
>>>>>>> d18076a (test: git tree state)

	return cmd
}
