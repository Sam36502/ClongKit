package lang

import (
	"errors"
	"fmt"
	"strings"
)

type Lexicon struct {
	Words []Word `json:"words"`
}

func (lang *Language) SetWord(word Word) error {
	_, i, err := lang.GetWord(word.GetRomanisation())
	if err == nil {
		// TODO: Merge new word? only pass updated values?
		lang.Lexicon.Words[i] = word
	} else {
		lang.Lexicon.Words = append(lang.Lexicon.Words, word)
	}

	return nil
}

func (lang *Language) DelWord(word Word) error {
	w, i, err := lang.GetWord(word.GetRomanisation())
	if err != nil {
		return errors.New(fmt.Sprintf("Word '%s' not found in the dictionary", w.GetRomanisation()))
	}
	lang.Lexicon.Words[i] = lang.Lexicon.Words[len(lang.Lexicon.Words)-1]
	lang.Lexicon.Words = lang.Lexicon.Words[:len(lang.Lexicon.Words)-1]
	return nil
}

// Gets a word from the dictionary and returns word, index, and error
func (lang *Language) GetWord(romanisation string) (*Word, uint, error) {
	for i, w := range lang.Lexicon.Words {
		if strings.ToLower(w.GetRomanisation()) == strings.ToLower(romanisation) {
			return &w, uint(i), nil
		}
	}
	return nil, 0, errors.New(fmt.Sprintf("No words with romanisation '%s' found.", romanisation))
}
