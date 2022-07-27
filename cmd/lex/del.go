/*
Copyright © 2022 Samuel Pearce

*/
package lex

import (
	"fmt"

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
		fmt.Println("del called")
	},
}

func init() {
	// Del Flags:
}
