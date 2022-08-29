package jsonfile

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Sam36502/ClongKit/presenter/lang"
)

func (sr *SyllableRule) String() string {
	bits := append(sr.OnsetGroups, sr.NucleusGroup)
	bits = append(bits, sr.CodaGroups...)
	return strings.Join(bits, "")
}

func (fls *JSONFileLangStorage) AddSyllableRule(patt lang.SyllableRule) error {
	_, err := fls.getSyllableRuleIndex(patt)
	if err != nil {
		jfp := SyllableRule{
			OnsetGroups:  patt.OnsetGroups,
			NucleusGroup: patt.NucleusGroup,
			CodaGroups:   patt.CodaGroups,
		}
		fls.lang.Phonology.SyllableRules = append(fls.lang.Phonology.SyllableRules, jfp)
		return nil
	}
	return errors.New(fmt.Sprintf("Pattern '%s' already exists", patt))
}

func (fls *JSONFileLangStorage) GetAllSyllableRules() ([]lang.SyllableRule, error) {
	arr := make([]lang.SyllableRule, len(fls.lang.Phonology.SyllableRules))
	for i, r := range fls.lang.Phonology.SyllableRules {
		arr[i] = r.toLang()
	}
	return arr, nil
}

func (fls *JSONFileLangStorage) DelSyllableRule(patt lang.SyllableRule) error {
	ind, err := fls.getSyllableRuleIndex(patt)
	if err != nil {
		return err
	}
	fls.lang.Phonology.SyllableRules = append(
		fls.lang.Phonology.SyllableRules[:ind],
		fls.lang.Phonology.SyllableRules[ind+1:]...,
	)
	return nil
}

func (fls *JSONFileLangStorage) getSyllableRuleIndex(findr lang.SyllableRule) (int, error) {
	for i, sr := range fls.lang.Phonology.SyllableRules {
		if strings.Compare(sr.String(), findr.String()) == 0 {
			return i, nil
		}
	}
	return -1, errors.New(fmt.Sprintf("No syllable rule '%s' found", findr.String()))
}
