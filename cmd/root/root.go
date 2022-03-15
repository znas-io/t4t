package root

import (
	"github.com/spf13/cobra"
	"github.com/znas-io/t4t/cmd/add"
	"github.com/znas-io/t4t/cmd/get"
)

const (
	use   = "t4t"
	short = "Helps organize filesystem using tags"
	long  = ``
)

func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
	}

	c.AddCommand(add.NewCommand())
	c.AddCommand(get.NewCommand())

	return c
}
