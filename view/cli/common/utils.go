package common

import (
	"errors"
	"fmt"

	"github.com/Sam36502/ClongKit/model"
	"github.com/Sam36502/ClongKit/model/jsonfile"
	"github.com/spf13/cobra"
)

// Common utility to get language file
func GetLang(cbr *cobra.Command) (model.LangStorage, error) {

	// TODO: Add other options for lang storage (MongoDB, SQL, etc.)
	langFileFlag := cbr.Flag(LangFileFlag)
	indentFlag := cbr.Flag(IndentFlag)
	store, err := jsonfile.NewFileLangStorage(langFileFlag.Value.String(), indentFlag.Changed)
	if err != nil {
		// Try to create new file
		store := &jsonfile.JSONFileLangStorage{}
		store.SetFilename(langFileFlag.Value.String())
		err := store.Close()
		if err != nil {
			msg := fmt.Sprintf("Failed to create language file '%s': %s\n", langFileFlag.Value.String(), err.Error())
			return nil, errors.New(msg)
		}

		store, err = jsonfile.NewFileLangStorage(langFileFlag.Value.String(), indentFlag.Changed)
		if err != nil {
			msg := fmt.Sprintf("Failed to load language file '%s': %s\n", langFileFlag.Value.String(), err.Error())
			return nil, errors.New(msg)
		}
		return store, nil
	}
	return store, nil
}
