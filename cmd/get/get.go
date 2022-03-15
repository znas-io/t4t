package get

import (
	"github.com/spf13/cobra"
)

const (
	use   = "get"
	short = "Gets paths that match all the provided tags"
	long  = ``
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run:   run,
	}
}

func run(*cobra.Command, []string) {
}
