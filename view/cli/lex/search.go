/*
Copyright © 2022 Samuel Pearce

*/
package lex

import (
	"fmt"
	"strings"

	"github.com/Sam36502/ClongKit/view/cli/common"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches the dictionary for a word",
	Long: `Takes a romanisation, meaning, or other filter
and searches the dictionary for all words matching the query.
If a romanisation is used (because they must be unique), it
will return one result with all the information it has.`,
	Run: func(cmd *cobra.Command, args []string) {
		langstore, err := common.GetLang(cmd)
		if err != nil {
			fmt.Printf("Failed to initialise language storage:\n%s\n", err.Error())
			return
		}

		wrds, err := langstore.GetAllWords()
		if err != nil {
			fmt.Printf("Failed to fetch all words:\n%s\n", err.Error())
			return
		}
		for _, w := range wrds {
			fmt.Println("Word:", w.GetRomanisation())
			fmt.Println("Pronunciation:", w.GetPronunciation())
			fmt.Println("Meanings:", strings.Join(w.Meanings, ", "))
			fmt.Print("\n")
		}
	},
}

func init() {
	// Search Flags:
}
