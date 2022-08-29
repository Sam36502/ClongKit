/*
Copyright © 2022 Samuel Pearce
*/
package tactics

import (
	"fmt"

	"github.com/Sam36502/ClongKit/view/cli/common"
	"github.com/spf13/cobra"
)

var TacticsCmd = &cobra.Command{
	Use:   "tactics",
	Short: "Commands for adding, listing and deleting syllable patterns",
	Long:  `Commands to add, list and delete syllable patterns for generating words from the phonology.`,
	Run: func(cmd *cobra.Command, args []string) {
		langstore, err := common.GetLang(cmd)
		if err != nil {
			fmt.Printf("Failed to initialise language storage:\n%s\n", err.Error())
			return
		}
		patterns, err := langstore.GetAllSyllableRules()
		if err != nil {
			fmt.Println("Failed to fetch syllable patterns")
			return
		}

		if len(patterns) == 0 {
			fmt.Println("No syllable patterns set")
			return
		}

		for _, p := range patterns {
			fmt.Printf("    %s\n", p.String())
		}
	},
}

func init() {
	// Tactics Flags:

	// Tactics Subcommands:
	TacticsCmd.AddCommand(AddCmd)
	TacticsCmd.AddCommand(DelCmd)

}
