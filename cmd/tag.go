package cmd

/*
Copyright © 2022 Jose Sanz <znas@znas.io>
*/
import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/znas-io/t4t/core"
	"strings"
)

const (
	short         = "Tags one or more files and directories to one of more tags"
	long          = ``
	validTagRegex = "^[a-zA-Z0-9-]+$"
)

// tagCmd represents the tag command
var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: short,
	Long:  long,
	Args:  cobra.MatchAll(cobra.MinimumNArgs(2), validateArgs),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test")
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tagCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	tagCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//tagCmd.MarkFlagCustom()
}

func validateArgs(_ *cobra.Command, args []string) error {
	tags := strings.Split(args[0], ",")

	var err error

	if err = core.ValidateTags(validTagRegex, tags...); err != nil {
		return err
	}

	if err = core.ValidatePaths(args[1:]...); err != nil {
		return err
	}

	return nil
}
