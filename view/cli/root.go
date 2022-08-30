/*
Copyright © 2022 Samuel Pearce
*/
package cli

import (
	"os"

	"github.com/Sam36502/ClongKit/model/jsonfile"
	"github.com/Sam36502/ClongKit/view/cli/common"
	"github.com/Sam36502/ClongKit/view/cli/lex"
	"github.com/Sam36502/ClongKit/view/cli/phono"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "clongkit",
	Short: "A tool for managing and using conlangs",
	Long: `This tool is meant to aid in the
use and generation of constructed languages for quick
use. In particular, it was made for inventing phrases
on-the-fly during TTRPG games.

It can also simply be used as a fast and efficient tool
for keeping all your constructed languages, or natural
languages, organised.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Global Flags
	rootCmd.PersistentFlags().StringP(common.LangFileFlag, "f", jsonfile.DefaultLanguageFile, "The file to load your language from")
	rootCmd.PersistentFlags().BoolP(common.IndentFlag, "i", false, "Keeps lang-file format readable")

	// Add Commands
	rootCmd.AddCommand(lex.LexCmd)
	rootCmd.AddCommand(phono.PhonoCmd)
}
