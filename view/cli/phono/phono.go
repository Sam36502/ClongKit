/*
Copyright © 2022 Samuel Pearce

*/
package phono

import (
	"github.com/spf13/cobra"
)

var PhonoCmd = &cobra.Command{
	Use:   "phono",
	Short: "All commands for working with your language's sounds",
	Long: `All commands for generating, building, analysing and using
your phonology & phonotactics as well as orthography & romanisation.`,
	/*
		Run: func(cmd *cobra.Command, args []string) {
			_, err := common.GetLang(cmd)
			if err != nil {
				fmt.Printf("Failed to initialise language storage:\n%s\n", err.Error())
				return
			}

			fmt.Println("phono called!\nThis is a test.")
		},
	*/
}

func init() {
	// Phono Flags:

	// Phono Subcommands:
	PhonoCmd.AddCommand(SetCmd)
	PhonoCmd.AddCommand(GetCmd)
	PhonoCmd.AddCommand(DelCmd)

}
