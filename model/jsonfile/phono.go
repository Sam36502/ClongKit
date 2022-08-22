package jsonfile

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Sam36502/ClongKit/presenter/lang"
)

func (fls *JSONFileLangStorage) SetPhoneme(ph lang.Phoneme) error {
	gp, i, err := fls.getPhonemeIndex(ph.Romanisation)
	if err != nil {
		fls.lang.Phonology.Phonemes = append(fls.lang.Phonology.Phonemes, Phoneme(*gp))
	}
	fls.lang.Phonology.Phonemes[i] = Phoneme(ph)
	return nil
}

func (fls *JSONFileLangStorage) GetPhoneme(rom string) (*lang.Phoneme, error) {
	ph, _, err := fls.getPhonemeIndex(rom)
	return ph, err
}

func (fls *JSONFileLangStorage) getPhonemeIndex(rom string) (*lang.Phoneme, uint64, error) {
	for i, ph := range fls.lang.Phonology.Phonemes {
		if strings.Compare(strings.ToLower(ph.Romanisation), strings.ToLower(rom)) == 0 {
			lph := lang.Phoneme(ph)
			return &lph, uint64(i), nil
		}
	}
	return nil, 0, errors.New(fmt.Sprintf("No phoneme found with romanisation '%s'", rom))
}

func (fls *JSONFileLangStorage) GetAllPhonemes() ([]lang.Phoneme, error) {
	phs := make([]lang.Phoneme, len(fls.lang.Phonology.Phonemes))
	for i, p := range fls.lang.Phonology.Phonemes {
		phs[i] = lang.Phoneme(p)
	}
	return phs, nil
}

func (fls *JSONFileLangStorage) DelPhoneme(rom string) error {
	_, i, err := fls.getPhonemeIndex(rom)
	if err != nil {
		return err
	}

	fls.lang.Phonology.Phonemes = append(fls.lang.Phonology.Phonemes[:i], fls.lang.Phonology.Phonemes[i+1:]...)
	return nil
}

func (fls *JSONFileLangStorage) AddSyllablePattern(patt string) error {
	for _, r := range fls.lang.Phonology.Phonotactics.Rules {
		if strings.Compare(strings.ToLower(r), strings.ToLower(patt)) == 0 {
			return errors.New(fmt.Sprintf("phonotactic rule '%s' already exists", patt))
		}
	}
	fls.lang.Phonology.Phonotactics.Rules = append(fls.lang.Phonology.Phonotactics.Rules, patt)
	return nil
}

func (fls *JSONFileLangStorage) GetAllSyllablePatterns() ([]string, error) {
	return fls.lang.Phonology.Phonotactics.Rules, nil
}

func (fls *JSONFileLangStorage) DelSyllablePattern(patt string) error {
	rules := fls.lang.Phonology.Phonotactics.Rules
	for i, r := range rules {
		if strings.Compare(strings.ToLower(r), strings.ToLower(patt)) == 0 {
			fls.lang.Phonology.Phonotactics.Rules = append(rules[:i], rules[i+1:]...)
			return nil
		}
	}
	return errors.New(fmt.Sprintf("No syllable rule '%s' found", patt))
}
