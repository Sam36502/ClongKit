package jsonfile

import "github.com/Sam36502/ClongKit/model"

var _ model.LangStorage = (*JSONFileLangStorage)(nil)

type Language struct {
	Name      string    `json:"name"`
	ID        string    `json:"id"`
	Phonology Phonology `json:"phonology"`
	Lexicon   Lexicon   `json:"lexicon"`
}

type Phonology struct {
	Phonemes      []Phoneme      `json:"phonemes"`
	SyllableRules []SyllableRule `json:"syllable_rules"`
}

type SyllableRule struct {
	OnsetGroups  []string `json:"ons"`
	NucleusGroup string   `json:"nuc"`
	CodaGroups   []string `json:"cod"`
}

type Phoneme struct {
	IPA          string   `json:"ipa"`
	Romanisation string   `json:"rom"`
	Groups       []string `json:"grp"`
}

type Lexicon struct {
	Words []Word `json:"words"`
}

type Word struct {
	Romanisation string   `json:"rom"`
	Etymology    string   `json:"ety"`
	Meanings     []string `json:"mns"`
	Tags         []string `json:"tgs"`
}

func (fls *JSONFileLangStorage) GetName() (string, error) {
	return fls.lang.Name, nil
}

func (fls *JSONFileLangStorage) GetID() (string, error) {
	return fls.lang.ID, nil
}

func (fls *JSONFileLangStorage) SetName(name string) error {
	fls.lang.Name = name
	return nil
}

func (fls *JSONFileLangStorage) SetID(ID string) error {
	fls.lang.ID = ID
	return nil
}
