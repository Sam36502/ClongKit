/*
Copyright © 2022 Samuel Pearce
*/
package lex

import (
	"github.com/spf13/cobra"
)

// lexCmd represents the lex command
var LexCmd = &cobra.Command{
	Use:   "lex",
	Short: "All the commands to do with storing and searching words",
	Long: `All the commands that handle the storing and searching
of your language's lexemese (words)`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("lex called")
	// },
}

func init() {
	LexCmd.AddCommand(setCmd)
	LexCmd.AddCommand(searchCmd)
	LexCmd.AddCommand(delCmd)
	LexCmd.AddCommand(genCmd)
}
