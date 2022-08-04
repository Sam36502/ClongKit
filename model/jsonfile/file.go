package jsonfile

import (
	"encoding/json"
	"io/ioutil"
)

const (
	DefaultLanguageFile = "lang.json"
	FileIndentation     = "    " // Four spaces
	LangfilePermissions = 0644
)

type JSONFileLangStorage struct {
	filename string
	lang     Language
	indented bool
}

func NewFileLangStorage(filename string, indentedFormatting bool) (*JSONFileLangStorage, error) {
	fls := JSONFileLangStorage{
		filename: filename,
		indented: indentedFormatting,
	}
	l, err := fls.loadLanguage()
	if err != nil {
		return nil, err
	}
	fls.lang = *l
	return &fls, nil
}

func (fls *JSONFileLangStorage) Close() error {
	return fls.saveLanguage()
}

func (fls *JSONFileLangStorage) loadLanguage() (*Language, error) {

	// Read file
	data, err := ioutil.ReadFile(fls.filename)
	if err != nil {
		return nil, err
	}

	// Parse file
	var jsonLang Language
	err = json.Unmarshal(data, &jsonLang)
	if err != nil {
		return nil, err
	}

	return &jsonLang, nil
}

func (fls *JSONFileLangStorage) saveLanguage() error {

	// Marshal data
	var data []byte
	var err error
	if err != nil {
		return err
	}

	if fls.indented {
		data, err = json.MarshalIndent(fls.lang, "", FileIndentation)
	} else {
		data, err = json.Marshal(fls.lang)
	}

	if err != nil {
		return err
	}

	// Save file
	err = ioutil.WriteFile(fls.filename, data, LangfilePermissions)
	if err != nil {
		return err
	}

	return nil
}
