package add

import (
	"github.com/spf13/cobra"
	"github.com/znas-io/t4t/core"
	"os"
)

const (
	use   = "add"
	short = "Adds tags to one or more files or directories"
	long  = ``
)

var (
	entries        = core.NewSortedEntriesMap()
	arrayTagsInput = make([]string, 0)
	sliceTagsInput = make([]string, 0)
)

func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Args:  cobra.MatchAll(cobra.MinimumNArgs(1), parseAndValidateInput),
		Run:   run,
	}

	c.Flags().StringArrayVarP(&arrayTagsInput, "tag", "t", make([]string, 0), "-t foo -t bar")
	c.Flags().StringSliceVar(&sliceTagsInput, "tags", make([]string, 0), "--tags foo,bar")

	return c
}

func run(*cobra.Command, []string) {
	var f *os.File
	var m map[string]*core.Entry
	var err error

	partition := ""

	for _, e := range entries.GetEntries() {
		p := e.GetTagPartition()

		if partition != p {
			partition = p

			if f != nil {
				_ = f.Close()
				cobra.CheckErr(err)
			}

			if f, err = core.GetDataFile(partition); err != nil {
				cobra.CheckErr(err)
			}

			if m, err = core.MapFileEntriesByID(f); err != nil {
				_ = f.Close()
				cobra.CheckErr(err)
			}
		}

		if _, ok := m[e.GetID()]; ok {
			continue
		}

		m[e.GetID()] = e

		if _, err = f.WriteString(e.String()); err != nil {
			_ = f.Close()
			cobra.CheckErr(err)
		}
	}

	err = f.Close()
	cobra.CheckErr(err)
}

func parseAndValidateInput(_ *cobra.Command, args []string) error {
	tags := append(arrayTagsInput, sliceTagsInput...)

	for _, tag := range tags {
		for _, path := range args {
			var i *core.Entry
			var err error

			if i, err = core.NewEntry(tag, path); err != nil {
				return err
			}

			if err = entries.Add(i); err != nil {
				return err
			}
		}
	}
	return nil
}
