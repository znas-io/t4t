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
	short         = "Adds tags to one or more files or directories"
	long          = ``
	validTagRegex = "^[a-zA-Z0-9-]+$"
)

var (
	arrayTags = make([]string, 0)
	sliceTags = make([]string, 0)

	addCmd = &cobra.Command{
		Use:   "add",
		Short: short,
		Long:  long,
		Args:  cobra.MatchAll(cobra.MinimumNArgs(1), validateArgs),
		Run: func(cmd *cobra.Command, args []string) {
			tags := inputTags()
			fmt.Println(tags)
		},
	}
)

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringArrayVarP(&arrayTags, "tag", "t", make([]string, 0), "-t foo -t bar")
	addCmd.Flags().StringSliceVar(&sliceTags, "tags", make([]string, 0), "--tags foo,bar")
}

func inputTags() []string {
	return append(arrayTags, sliceTags...)
}

func validateArgs(_ *cobra.Command, args []string) error {
	var err error

	if err = core.ValidateTags(validTagRegex, inputTags()...); err != nil {
		return err
	}

	if err = core.ValidatePaths(args...); err != nil {
		return err
	}

	return nil
}
