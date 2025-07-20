package cmd

import (
	"fmt"
	"os"

	"net/url"

	"github.com/spf13/cobra"
)

var batch string
var single string

func isValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}

// ripCmd represents the rip command
var ripCmd = &cobra.Command{
	Use:   "rip",
	Short: "The rip command downloads full-length Panopto recordings based on a recording URL",
	RunE: func(cmd *cobra.Command, args []string) error {
		if single != "" {
			if !isValidURL(single) {
				return fmt.Errorf("invalid URL: %s", single)
			}
			// process single URL...
		}

		if batch != "" {
			if _, err := os.Stat(batch); err != nil {
				return fmt.Errorf("cannot read file: %s", batch)
			}
			// process batch file...
		}

		return nil
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if (single == "" && batch == "") || (single != "" && batch != "") {
			cmd.Help()
			return fmt.Errorf("you must specify exactly one of --single or --batch")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(ripCmd)

	ripCmd.Flags().StringVarP(&batch, "batch", "b", "", "Path to a file containing URLs to download")
	ripCmd.Flags().StringVarP(&single, "single", "s", "", "A single Panopto video URL to download")
}
