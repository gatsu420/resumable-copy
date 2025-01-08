package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "resumable-copy",
	Short: "Demonstration of copying files that can be resumed using byte index",
	Long:  `Demonstration of copying files that can be resumed using byte index`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(copyCmd)
}
