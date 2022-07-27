/*
Copyright © 2022 Samuel Pearce

*/
package lex

import (
	"fmt"
	"strings"

	"github.com/Sam36502/ClongKit/cmd/common"
	"github.com/Sam36502/ClongKit/internal/storage"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Adds or updates a word in your dictionary",
	Long: `takes a romanisation and list of meanings and
either adds them to the dictionary or updates an existing entry.

It's recommended to add as many meanings as you can think of
to improve the chance of finding it when searching the dictionry.`,
	Run: func(cmd *cobra.Command, args []string) {
		lang, filename := common.GetLang(cmd)
		if lang == nil {
			return
		}

		if len(args) != 1 {
			fmt.Println("Exactly one argument is required: the word romanisation")
			return
		}
		rom := args[0]

		// Parse romanisation
		word, err := lang.ParseWord(rom)
		if err != nil || word == nil {
			fmt.Println("Failed to parse the romanisation, are all the letters registered in the phonology?")
			return
		}

		// Add meanings
		means := cmd.Flag(common.MeaningListFlag)
		if means.Changed {
			word.Meanings = strings.Split(means.Value.String(), common.ListSeparator)
			for i, m := range word.Meanings {
				word.Meanings[i] = strings.TrimSpace(m)
			}
		}

		// TODO: Look into custom JSON (un)marshaler (Less redundant info)
		lang.SetWord(*word)

		// Save changes to file
		err = storage.SaveLanguage(lang, filename, cmd.Flag(common.PrettyFileFlag).Changed)
		if err != nil {
			fmt.Printf("Failed to save language '%s': %s\n", filename, err.Error())
			return
		}
	},
}

func init() {
	// Set Flags:
	setCmd.Flags().StringP(common.MeaningListFlag, "m", "", "List of meanings; what the word means in your mother-tongue.")
}
