package cmd

/*
Copyright © 2022 Jose Sanz <znas@znas.io>
*/
import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/znas-io/t4t/core"
)

const (
	short = "Adds tags to one or more files or directories"
	long  = ``
)

var (
	addCmd = &cobra.Command{
		Use:   "add",
		Short: short,
		Long:  long,
		Args:  cobra.MatchAll(cobra.MinimumNArgs(1), parseAndValidateInput),
		Run:   run,
	}
)

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringArrayVarP(&arrayTagsInput, "tag", "t", make([]string, 0), "-t foo -t bar")
	addCmd.Flags().StringSliceVar(&sliceTagsInput, "tags", make([]string, 0), "--tags foo,bar")
}

func run(*cobra.Command, []string) {
	for _, i := range sortedTagsMap.GetEntries() {
		fmt.Println(i.FileString())
	}
}

func parseAndValidateInput(_ *cobra.Command, args []string) error {
	tags := append(arrayTagsInput, sliceTagsInput...)

	for _, tag := range tags {
		for _, path := range args {
			var i *core.Entry
			var err error

			if i, err = core.NewTag(tag, path); err != nil {
				return err
			}

			if err = sortedTagsMap.Add(i); err != nil {
				return err
			}
		}
	}
	return nil
}
