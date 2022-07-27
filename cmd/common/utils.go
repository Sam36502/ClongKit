package common

import (
	"fmt"

	"github.com/Sam36502/ClongKit/internal/lang"
	"github.com/Sam36502/ClongKit/internal/storage"
	"github.com/spf13/cobra"
)

// Common utility to get language file
func GetLang(cbr *cobra.Command) (*lang.Language, string) {
	langFileFlag := cbr.Flag(LangFileFlag)
	lang, err := storage.LoadLanguage(langFileFlag.Value.String())
	if err != nil {
		fmt.Printf("Failed to load language '%s': %s\n", langFileFlag.Value.String(), err.Error())
		return nil, ""
	}
	return lang, langFileFlag.Value.String()
}
