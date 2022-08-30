/*
Copyright © 2022 Samuel Pearce
*/
package phono

import (
	"fmt"
	"sort"

	"github.com/Sam36502/ClongKit/presenter/lang"
	"github.com/Sam36502/ClongKit/view/cli/common"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a list of phonemes in the language",
	Long: `Gets a list of phonemes in the language
broken down by group (Consonant, Vowel, etc.)
It can also be filtered with some basic options.`,
	Run: func(cmd *cobra.Command, args []string) {
		langstore, err := common.GetLang(cmd)
		if err != nil {
			fmt.Printf("Failed to initialise language storage:\n%s\n", err.Error())
			return
		}

		phs, err := langstore.GetAllPhonemes()
		if err != nil {
			fmt.Println("Failed to get phonemes:", err)
			return
		}

		grps := map[string][]lang.Phoneme{}
		grpInds := []string{}
		for _, p := range phs {
			for _, g := range p.Groups {
				gr, e := grps[g]
				grps[g] = append(gr, p)
				if !e {
					grpInds = append(grpInds, g)
				}
			}
		}
		sort.Strings(grpInds)
		for _, g := range grpInds {
			fmt.Printf("Group '%s'\n", g)
			for _, p := range grps[g] {
				fmt.Printf("%s<%s> /%s/; ", common.OutputIndent, p.Romanisation, p.IPA)
			}
			fmt.Print("\n")
		}
	},
}

func init() {
	// Set Flags:
	//TODO: Add flags to filter by (mainly by group, I suppose)
	//GetCmd.Flags().StringP(common.IPAFlag, "p", "", "Phoneme pronunciation in IPA")
}
