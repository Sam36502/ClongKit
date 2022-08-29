/*
Copyright © 2022 Samuel Pearce
*/
package tactics

import (
	"fmt"
	"strings"

	"github.com/Sam36502/ClongKit/presenter/lang"
	"github.com/Sam36502/ClongKit/view/cli/common"
	"github.com/spf13/cobra"
)

var DelCmd = &cobra.Command{
	Use:   "del",
	Short: "Deletes a syllable pattern from the language",
	Long:  `Deletes a syllable pattern from the language`,
	Run: func(cmd *cobra.Command, args []string) {
		langstore, err := common.GetLang(cmd)
		if err != nil {
			fmt.Printf("Failed to initialise language storage:\n%s\n", err.Error())
			return
		}

		if len(args) != 1 {
			fmt.Println("Exactly one argument is required: syllable pattern (C,C;V;C,N)")
			return
		}

		// Parse syllable rule
		arr := strings.Split(args[0], common.SyllableSeparator)
		rule := lang.SyllableRule{
			OnsetGroups:  strings.Split(arr[0], common.ListSeparator),
			NucleusGroup: arr[1],
			CodaGroups:   strings.Split(arr[2], common.ListSeparator),
		}

		phs, err := langstore.GetAllPhonemes()
		if err != nil {
			fmt.Println("Failed to get phonemes:", err)
		}
		err = rule.Validate(phs)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = langstore.DelSyllableRule(rule)
		if err != nil {
			fmt.Println("Failed to delete syllable pattern")
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
	// Del Flags:
}
