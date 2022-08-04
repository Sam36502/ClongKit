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

		// Get Search Params
		rom := ""
		romFlg := cmd.Flag(common.RomanisationFlag)
		if romFlg != nil && romFlg.Changed {
			rom = romFlg.Value.String()
		}

		etym := ""
		etymFlg := cmd.Flag(common.EtymologyFlag)
		if etymFlg != nil && etymFlg.Changed {
			etym = etymFlg.Value.String()
		}

		means := []string{}
		meansFlg := cmd.Flag(common.MeaningListFlag)
		if meansFlg != nil && meansFlg.Changed {
			means = strings.Split(meansFlg.Value.String(), common.ListSeparator)
		}

		tags := []string{}
		tagsFlg := cmd.Flag(common.TagListFlag)
		if tagsFlg != nil && tagsFlg.Changed {
			tags = strings.Split(tagsFlg.Value.String(), common.ListSeparator)
		}

		wrds, err := langstore.SearchWord(rom, etym, means, tags)
		if err != nil {
			fmt.Printf("Failed to search words:\n%s\n", err.Error())
			return
		}
		if len(wrds) == 0 {
			fmt.Println("No Words found")
		}
		for _, w := range wrds {
			fmt.Println("Word:", w.GetRomanisation())
			fmt.Printf("Pronunciation: /%s/\n", w.GetPronunciation())
			fmt.Println("Etymology:", w.Etymology)
			fmt.Println("Meanings:", strings.Join(w.Meanings, ", "))
			fmt.Println("Tags:", strings.Join(w.Tags, ", "))
			fmt.Print("\n")
		}
	},
}

func init() {
	// Search Flags:
	searchCmd.Flags().StringP(common.RomanisationFlag, "r", "", "The romanisation to search by (fuzzy)")
	searchCmd.Flags().StringP(common.EtymologyFlag, "e", "", "The etymology to search by (fuzzy)")
	searchCmd.Flags().StringP(common.MeaningListFlag, "m", "", "The meanings to search by (fuzzy)")
	searchCmd.Flags().StringP(common.TagListFlag, "t", "", "The tags to search by (exact)")
}
