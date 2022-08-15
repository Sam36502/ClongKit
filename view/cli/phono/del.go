/*
Copyright © 2022 Samuel Pearce

*/
package phono

import (
	"fmt"

	"github.com/Sam36502/ClongKit/view/cli/common"
	"github.com/spf13/cobra"
)

var DelCmd = &cobra.Command{
	Use:   "del",
	Short: "Removes a phoneme from the language",
	Long: `Takes the romanisation of a phoneme in the language
and removes it from the phonology.`,
	Run: func(cmd *cobra.Command, args []string) {
		langstore, err := common.GetLang(cmd)
		if err != nil {
			fmt.Printf("Failed to initialise language storage:\n%s\n", err.Error())
			return
		}

		if len(args) != 1 {
			fmt.Println("Exactly one argument is required: phoneme romanisation")
			return
		}
		ph, err := langstore.GetPhoneme(args[0])
		if err != nil {
			fmt.Printf("Couldn't find any phoneme with the romanisation '%s'\n", args[0])
			return
		}

		err = langstore.DelPhoneme(ph.Romanisation)
		if err != nil {
			fmt.Println("Failed to delete phoneme:", err)
			return
		}

		err = langstore.Close()
		if err != nil {
			fmt.Println("Failed to close language storage:", err)
			return
		}
	},
}

func init() {
	// Set Flags:
}
