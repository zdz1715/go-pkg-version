package go_pkg_version

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
)

const (
	versionShort   = `Print the version information.`
	versionLong    = `Print the version information.`
	versionExample = `  # Print the version information
  %s version
  # Print the version number only
  %s version -n 
  %s version --number`
)

type options struct {
	onlyNumber bool
}

// NewVersionCommand prints out the release version info for this command binary.
// It is used as a subcommand of a parent command.
func NewVersionCommand(out io.Writer, parentCommand string) *cobra.Command {
	o := &options{}
	cmd := &cobra.Command{
		Use:     "version",
		Short:   versionShort,
		Long:    versionLong,
		Example: fmt.Sprintf(versionExample, parentCommand, parentCommand, parentCommand),
		Run: func(cmd *cobra.Command, args []string) {
			_, _ = fmt.Fprintln(out, formatVersion(parentCommand, o.onlyNumber))
		},
	}
	cmd.Flags().BoolVarP(&o.onlyNumber, "number", "n", false, "Print the version number only")
	return cmd
}

func formatVersion(name string, onlyVersion bool) string {
	v := Get()
	if onlyVersion {
		return v.Version
	}
	return fmt.Sprintf("%#v", v)
}
