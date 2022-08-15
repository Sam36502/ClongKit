package lang

import (
	"errors"
	"fmt"
	"strings"
)

type Phoneme struct {
	IPA          string
	Romanisation string
	Groups       []string
}

type Phonology struct {
	Phonemes []Phoneme `json:"phonemes"`
}

func (ph *Phonology) ParseWord(rom string) (*Word, error) {
	newWord := Word{
		Phonemes: []Phoneme{},
	}
	for len(rom) > 0 {
		p, r, err := ph.parseOutFirstLetter(rom)
		if err != nil {
			return nil, err
		}
		rom = r
		newWord.Phonemes = append(newWord.Phonemes, *p)
	}
	return &newWord, nil
}

// Takes a romanisation and parses out the first letter it finds
// Returns the phoneme, the shortened string and a possible error
func (ph *Phonology) parseOutFirstLetter(romanisation string) (*Phoneme, string, error) {

	for i := len(romanisation); i >= 0; i-- {
		p, err := ph.ParsePhoneme(romanisation[:i])
		if err == nil {
			return p, romanisation[i:], nil
		}
	}

	return nil, romanisation, errors.New(fmt.Sprintf("No letter found within word string '%s'", romanisation))
}

// Takes a romanisation letter/cluster and returns the phoneme
// associated with that letter
func (ph *Phonology) ParsePhoneme(letter string) (*Phoneme, error) {
	for _, p := range ph.Phonemes {
		if strings.Compare(strings.ToLower(p.Romanisation), strings.ToLower(letter)) == 0 {
			return &p, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Failed to parse letter '%s'", letter))
}
