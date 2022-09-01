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

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a syllable pattern",
	Long:  `Adds a syllable pattern to the list of accepted ones for word generation.`,
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
		rule := lang.SyllableRule{}
		arr := strings.Split(args[0], common.SyllableSeparator)

		// Onset
		if len(arr) > 0 {
			gs := strings.Split(arr[0], common.ListSeparator)
			if len(gs) == 1 && gs[0] == "" {
				rule.OnsetGroups = []string{}
			} else {
				rule.OnsetGroups = gs
			}
		} else {
			rule.OnsetGroups = []string{}
		}

		// Nucleus
		if len(arr) > 1 {
			rule.NucleusGroup = arr[1]
		} else {
			rule.NucleusGroup = ""
		}

		// Coda
		if len(arr) > 2 {
			gs := strings.Split(arr[2], common.ListSeparator)
			if len(gs) == 1 && gs[0] == "" {
				rule.CodaGroups = []string{}
			} else {
				rule.CodaGroups = gs
			}
		} else {
			rule.CodaGroups = []string{}
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

		err = langstore.AddSyllableRule(rule)
		if err != nil {
			fmt.Println("Failed to add syllable pattern:", err)
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
	// Add Flags:
}
