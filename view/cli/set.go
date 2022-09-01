/*
Copyright © 2022 Samuel Pearce
*/
package cli

import (
	"fmt"
	"github.com/Sam36502/ClongKit/view/cli/common"
	"github.com/spf13/cobra"
)

var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Creates or updates a language's information",
	Long: `Takes a language name & ID and either creates a new file for it,
or updates an existing language's information.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Parse input
		if len(args) != 2 {
			fmt.Println("Exactly 2 arguments required: Name and ID")
			return
		}
		name := args[0]
		ID := args[1]

		langstore, err := common.GetLang(cmd)
		if err != nil {
			fmt.Println("Failed to open language storage:", err)
			return
		}

		langstore.SetName(name)
		langstore.SetID(ID)

		fmt.Printf("Updated language '%s' information.\n", name)
	},
}

func init() {
	// Set Flags
}
