package lang

import (
	"errors"
	"fmt"
	"strings"
)

type Word struct {
	Phonemes  []Phoneme
	Meanings  []string
	Etymology string
	Tags      []string
}

// TODO: Find more efficient way of parsing words
func (lang *Language) ParseWord(romanisation string) (*Word, error) {
	newWord := Word{
		Phonemes: []Phoneme{},
	}
	for len(romanisation) > 0 {
		p, r, err := lang.parseOutFirstLetter(romanisation)
		if err != nil {
			return nil, err
		}
		romanisation = r
		newWord.Phonemes = append(newWord.Phonemes, *p)
	}
	return &newWord, nil
}

func (lang *Language) parseOutFirstLetter(romanisation string) (*Phoneme, string, error) {
	for l := range romanisation {
		letter := ""
		sub := romanisation[l:]
		for i, r := range sub {
			letter = fmt.Sprint(letter, string(r))
			p, err := lang.ParsePhoneme(letter)
			if err == nil {
				return p, romanisation[l+i+1:], nil
			}
		}
	}
	return nil, romanisation, errors.New(fmt.Sprintf("No letter found within word string '%s'", romanisation))
}

// Takes a romanisation letter/cluster and returns the phoneme
// associated with that letter
func (lang *Language) ParsePhoneme(letter string) (*Phoneme, error) {
	for _, p := range lang.Phonology.Phonemes {
		if strings.ToLower(p.Romanisation) == strings.ToLower(letter) {
			return &p, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Failed to parse letter '%s'", letter))
}

func (w *Word) GetRomanisation() string {
	rom := strings.Builder{}
	for _, p := range w.Phonemes {
		rom.WriteString(p.Romanisation)
	}
	return rom.String()
}

func (w *Word) GetPronunciation() string {
	pron := strings.Builder{}
	for _, p := range w.Phonemes {
		pron.WriteString(p.IPA)
	}
	return pron.String()
}
