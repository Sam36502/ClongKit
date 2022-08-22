package jsonfile

import (
	"errors"
	"fmt"
	"strings"
)

func (sr *SyllableRule) String() string {
	bits := append(sr.OnsetGroups, sr.NucleusGroup)
	bits = append(bits, sr.CodaGroups...)
	return strings.Join(bits, "")
}

func (fls *JSONFileLangStorage) AddSyllableRule(patt lang.SyllableRule) error {
	fls.lang.Phonology.SyllableRules
}

func (fls *JSONFileLangStorage) GetAllSyllableRules() ([]lang.SyllableRule, error) {
}

func (fls *JSONFileLangStorage) DelSyllableRule(patt lang.SyllableRule) error {
}

func (fls *JSONFileLangStorage) getSyllableRuleIndex(findr SyllableRule) (int, error) {
	for i, sr := range fls.lang.Phonology.SyllableRules {
		if strings.Compare(sr.String(), findr.String()) == 0 {
			return i, nil
		}
	}
	return -1, errors.New(fmt.Sprintf("No syllable rule '%s' found", findr.String()))
}
