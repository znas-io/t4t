package get

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/znas-io/t4t/core"
	"os"
	"sort"
)

const (
	use   = "get"
	short = "Gets paths that match all the provided tags"
	long  = ``
)

var (
	tags = make([]string, 0)
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Args:  cobra.MatchAll(cobra.MinimumNArgs(1), validateTags),
		Run:   run,
	}
}

func run(*cobra.Command, []string) {
	var f *os.File
	var m map[string]core.Entries
	var err error

	partition := ""
	paths := make(map[string]int)

	for _, tag := range tags {
		p := core.GetTagPartition(tag)

		if partition != p {
			partition = p

			if f != nil {
				_ = f.Close()
				cobra.CheckErr(err)
			}

			if f, err = core.GetDataFile(partition); err != nil {
				cobra.CheckErr(err)
			}

			if m, err = core.MapFileEntriesByTag(f); err != nil {
				_ = f.Close()
				cobra.CheckErr(err)
			}
		}

		var entries core.Entries
		var ok bool

		if entries, ok = m[tag]; !ok {
			continue
		}

		for _, entry := range entries {
			paths[entry.GetPath()] = paths[entry.GetPath()] + 1
		}
	}

	err = f.Close()
	cobra.CheckErr(err)

	for path, n := range paths {
		if n != len(tags) {
			continue
		}

		fmt.Println(path)
	}
}

func validateTags(_ *cobra.Command, args []string) error {
	for _, tag := range args {
		var err error

		if tag, err = core.ValidateTag(tag); err != nil {
			return err
		}

		tags = append(tags, tag)
	}

	sort.Strings(tags)

	return nil
}
