/*
Copyright © 2022 Samuel Pearce

*/
package phono

import (
	"fmt"

	"github.com/Sam36502/ClongKit/internal/storage"
	"github.com/spf13/cobra"
)

var PhonoCmd = &cobra.Command{
	Use:   "phono",
	Short: "All commands for working with your language's sounds",
	Long: `All commands for generating, building, analysing and using
your phonology & phonotactics as well as orthography & romanisation.`,
	Run: func(cmd *cobra.Command, args []string) {
		lang, err := storage.LoadLanguage(storage.DefaultLanguageFile)
		if err != nil {
			fmt.Printf("Failed to load language file '%s'.\nPlease check the format.\n", storage.DefaultLanguageFile)
			fmt.Println(err.Error())
			return
		}

		fmt.Println("phono called\nTesting parsing")
		if len(args) < 1 {
			fmt.Println("Word argument required")
			return
		}

		word, err := lang.ParseWord(args[0])
		if err != nil {
			fmt.Printf("Failed to parse word '%s'.\nPlease check the format.\n", args[0])
			fmt.Println(err.Error())
			return
		}

		fmt.Printf("Romanisation: %*s\nPronunciation: %*s\n", 50, word.GetRomanisation(), 50, word.GetPronunciation())
	},
}

func init() {
	// Phono Flags:

	// Phono Subcommands:

}
