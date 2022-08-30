/*
Copyright © 2022 Samuel Pearce
*/
package lex

import (
	"fmt"

	"github.com/Sam36502/ClongKit/view/cli/common"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generates a word from your phonology and phonotactics",
	Long: `Generates words based on your phonemes and syllable patterns.

You can change the number of words and range of word length with -n and -l.`,
	Run: func(cmd *cobra.Command, args []string) {
		langstore, err := common.GetLang(cmd)
		if err != nil {
			fmt.Printf("Failed to initialise language storage:\n%s\n", err.Error())
			return
		}

		// Parse Number
		n, err := cmd.Flags().GetInt(common.NumberFlag)
		if err != nil {
			fmt.Println("Failed to parse number:", err)
			return
		}

		// Parse Range
		l, err := cmd.Flags().GetString(common.LengthFlag)
		if err != nil {
			fmt.Println("Failed to parse length range:", err)
			return
		}

		// Generate Words
		for i := 0; i < n; i++ {
			str, err := langstore.GenerateWord(l)
			if err != nil {
				fmt.Println("Failed to generate word:", err)
				return
			}
			wrd, err := langstore.ParseWord(str)
			if err != nil {
				fmt.Println("Failed to parse word:", err)
				return
			}
			fmt.Printf("%s'%s' /%s/\n", common.OutputIndent, wrd.GetRomanisation(), wrd.GetPronunciation())
		}
	},
}

func init() {
	// Gen Flags:
	genCmd.Flags().StringP(common.LengthFlag, "l", "1-4", "Length of words in syllables (default 1-4)")
	genCmd.Flags().IntP(common.NumberFlag, "n", 20, "Number of words to generate (default 20)")
}
