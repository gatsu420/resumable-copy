package cmd

import (
	"log"

	"github.com/gatsu420/resumable-copy/src"
	"github.com/spf13/cobra"
)

var (
	srcFile   string
	destFile  string
	resumeAt  int
	chunkSize int
	lag       int
)

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy file",
	Long: `Copy content using byte index of source file.
It loops based on chunk size. 
	`,

	Run: func(cmd *cobra.Command, args []string) {
		err := src.ResumableCopy(srcFile, destFile, resumeAt, chunkSize, lag)
		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	copyCmd.Flags().StringVar(&srcFile, "src", "", "Source file")
	copyCmd.Flags().StringVar(&destFile, "dest", "destination", "Destination file")
	copyCmd.Flags().IntVar(&resumeAt, "resume-at", 0, "Byte index to start copying from")
	copyCmd.Flags().IntVar(&chunkSize, "chunk-size", 4, "Byte chunk size per copy iteration")
	copyCmd.Flags().IntVar(&lag, "lag", 3, "Simulated lag in second while copying each chunk")

	copyCmd.MarkFlagRequired("source")
}
