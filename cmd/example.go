package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// dateCmd represents the date command
var exampleCmd = &cobra.Command{
	Use:   "example",
	Short: "Does nothing useful.",
	Run: func(cmd *cobra.Command, args []string) {
		letters := pterm.NewLettersFromString("useless")
		pterm.DefaultBigText.WithLetters(letters).Render()
	},
}

func init() {
	rootCmd.AddCommand(exampleCmd)

	exampleCmd.Flags().StringP("format", "f", "02 Jan 06", "specify a custom date format")
}
