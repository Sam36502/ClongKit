/*
Copyright © 2022 Samuel Pearce

*/
package lex

import (
	"fmt"

	"github.com/Sam36502/ClongKit/view/cli/common"
	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Takes a word's romanisation and removes it from the dictionary",
	Long: `Takes a word's romanisation and removes the entry
from the dictionary entirely. Warning! This cannot be undone, except by
manually adding the word and all its information again!`,
	Run: func(cbr *cobra.Command, args []string) {
		langstore, err := common.GetLang(cbr)
		if err != nil {
			fmt.Printf("Failed to initialise language storage:\n%s\n", err.Error())
			return
		}

		if len(args) != 1 {
			fmt.Println("Exactly one argument is required: the word romanisation")
			return
		}
		rom := args[0]

		err = langstore.DelWord(rom)
		if err != nil {
			fmt.Printf("Failed to delete '%s' from the dictionary\n", rom)
			return
		}

		err = langstore.Close()
		if err != nil {
			fmt.Println("Failed to close language storage: ", err)
			return
		}
	},
}

func init() {
	// Del Flags:
}
