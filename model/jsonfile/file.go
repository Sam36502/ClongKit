package jsonfile

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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

func (fls *JSONFileLangStorage) SetFilename(fn string) {
	fls.filename = fn
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

func (fls *JSONFileLangStorage) CreateLanguage(name, ID string) error {

	// Check file already exists
	_, err := os.Stat(fls.filename)
	if !os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("File '%s' already exists", fls.filename))
	}

	// Otherwise create new one
	fls.lang = Language{
		Name: name,
		ID:   ID,
	}

	err = fls.saveLanguage()
	if err != nil {
		return err
	}
	return nil
}
