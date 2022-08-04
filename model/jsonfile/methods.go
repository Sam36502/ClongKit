package jsonfile

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Sam36502/ClongKit/model"
	"github.com/Sam36502/ClongKit/presentor/lang"
)

var _ model.LangStorage = (*JSONFileLangStorage)(nil)

func (fls *JSONFileLangStorage) SetWord(wrd lang.Word) error {
	mergeWord, ind, err := fls.getWordIndex(wrd.GetRomanisation())
	if err != nil {
		fls.lang.Lexicon.Words = append(fls.lang.Lexicon.Words, Word{
			Romanisation: wrd.GetRomanisation(),
			Etymology:    wrd.Etymology,
			Meanings:     wrd.Meanings,
			Tags:         wrd.Tags,
		})
		return nil
	}
	if wrd.GetRomanisation() != "" {
		mergeWord.Romanisation = wrd.GetRomanisation()
	}
	if wrd.Etymology != "" {
		mergeWord.Etymology = wrd.Etymology
	}
	if len(wrd.Meanings) > 0 {
		mergeWord.Meanings = wrd.Meanings
	}
	if len(wrd.Tags) > 0 {
		mergeWord.Tags = wrd.Tags
	}
	fls.lang.Lexicon.Words[ind] = *mergeWord
	return nil
}

func (fls *JSONFileLangStorage) getWordIndex(rom string) (*Word, uint64, error) {
	lowRom := strings.ToLower(rom)
	for i, w := range fls.lang.Lexicon.Words {
		if strings.Compare(strings.ToLower(w.Romanisation), lowRom) == 0 {
			return &w, uint64(i), nil
		}
	}
	return nil, 0, errors.New(fmt.Sprintf("Could not find word with romanisation '%s'", rom))
}

func (fls *JSONFileLangStorage) GetWord(rom string) (*lang.Word, error) {
	jsonWrd, _, err := fls.getWordIndex(rom)
	if err != nil {
		return nil, err
	}
	return jsonWrd.toLang(&fls.lang.Phonology)
}

func (fls *JSONFileLangStorage) GetAllWords() ([]lang.Word, error) {
	wrds := make([]lang.Word, len(fls.lang.Lexicon.Words))
	for i, w := range fls.lang.Lexicon.Words {
		lw, err := w.toLang(&fls.lang.Phonology)
		if err != nil {
			return nil, err
		}
		wrds[i] = *lw
	}
	return wrds, nil
}

func (fls *JSONFileLangStorage) SearchWord(rom, pron, etym, means, tags string) ([]lang.Word, error) {
	return nil, nil
}

func (fls *JSONFileLangStorage) DelWord(rom string) error {
	return nil
}

func (fls *JSONFileLangStorage) ParseWord(rom string) (*lang.Word, error) {
	lp := fls.lang.Phonology.toLang()
	return lp.ParseWord(rom)
}
