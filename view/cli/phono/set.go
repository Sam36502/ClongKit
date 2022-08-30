/*
Copyright © 2022 Samuel Pearce
*/
package phono

import (
	"fmt"
	"strings"

	"github.com/Sam36502/ClongKit/presenter/lang"
	"github.com/Sam36502/ClongKit/view/cli/common"
	"github.com/spf13/cobra"
)

var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Adds or updates a phoneme in your phonology",
	Long: `Takes a phoneme romanisation and phoneme attributes
and either updates that phoneme with the new values, or adds it to
your phonology, if it doesn't already exist.`,
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
		ph := &lang.Phoneme{
			Romanisation: args[0],
			IPA:          "",
			Groups:       []string{},
		}

		ipaFlg := cmd.Flag(common.IPAFlag)
		if ipaFlg.Changed {
			ph.IPA = ipaFlg.Value.String()
		}

		grpsFlg := cmd.Flag(common.GroupsFlag)
		if grpsFlg.Changed {
			ph.Groups = strings.Split(grpsFlg.Value.String(), common.ListSeparator)
			for i, p := range ph.Groups {
				ph.Groups[i] = strings.TrimSpace(p)
			}
		}

		err = langstore.SetPhoneme(*ph)
		if err != nil {
			fmt.Println("Failed to set phoneme")
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
	SetCmd.Flags().StringP(common.IPAFlag, "p", "", "Phoneme pronunciation in IPA")
	SetCmd.Flags().StringP(common.GroupsFlag, "g", "", "Phoneme groups e.g. (C,N) -> (Consonant, Nasal)")
}
