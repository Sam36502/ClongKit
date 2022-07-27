package storage

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Sam36502/ClongKit/internal/lang"
)

const (
	DefaultLanguageFile = "lang.json"
	JSONFileIndentation = "    " // Four spaces
	LangfilePermissions = 0644
)

func LoadLanguage(filename string) (*lang.Language, error) {

	// Read file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Parse file
	var jsonLang JSONLanguage
	err = json.Unmarshal(data, &jsonLang)
	if err != nil {
		return nil, err
	}

	return jsonLang.ToLanguage()
}

func SaveLanguage(l *lang.Language, filename string, pretty bool) error {

	// Marshal data
	var data []byte
	var err error
	if pretty {
		data, err = json.MarshalIndent(l, "", JSONFileIndentation)
	} else {
		data, err = json.Marshal(l)
	}

	if err != nil {
		return err
	}

	// Save file
	err = ioutil.WriteFile(filename, data, LangfilePermissions)
	if err != nil {
		return err
	}

	// TODO: Convert to JSONlang before saving
	return nil
}
