package jsonfile

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Sam36502/ClongKit/lang"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

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

func (fls *JSONFileLangStorage) SearchWord(rom, etym string, means, tags []string) ([]lang.Word, error) {
	words := fls.lang.Lexicon.Words
	filtered := false

	if rom != "" {
		i := 0
		for _, w := range words {
			if fuzzy.Match(strings.ToLower(rom), strings.ToLower(w.Romanisation)) {
				words[i] = w
				i++
			}

		}
		filtered = true

		// Empty rest of array
		for j := i; j < len(words); j++ {
			words[j] = Word{}
		}
		words = words[:i]
	}

	if etym != "" {
		i := 0
		for _, w := range words {
			if fuzzy.Match(strings.ToLower(etym), strings.ToLower(w.Etymology)) {
				words[i] = w
				i++
			}
		}
		filtered = true

		// Empty rest of array
		for j := i; j < len(words); j++ {
			words[j] = Word{}
		}
		words = words[:i]
	}

	if len(means) > 0 {
		i := 0
		for _, w := range words {
			for _, wm := range w.Meanings {
				matched := false
				for _, im := range means {
					if fuzzy.Match(strings.ToLower(im), strings.ToLower(wm)) {
						words[i] = w
						i++
						matched = true
						break
					}
					if matched {
						break
					}
				}
			}
		}
		filtered = true

		// Empty rest of array
		for j := i; j < len(words); j++ {
			words[j] = Word{}
		}
		words = words[:i]
	}

	if len(tags) > 0 {
		i := 0
		for _, w := range words {
			for _, wt := range w.Tags {
				matched := false
				for _, it := range tags {
					if strings.Compare(strings.ToLower(wt), strings.ToLower(it)) == 0 {
						words[i] = w
						i++
						matched = true
						break
					}
				}
				if matched {
					break
				}
			}
		}
		filtered = true

		// Empty rest of array
		for j := i; j < len(words); j++ {
			words[j] = Word{}
		}
		words = words[:i]
	}

	if !filtered {
		return nil, nil
	}

	lwords := make([]lang.Word, len(words))
	for i, w := range words {
		lw, err := w.toLang(&fls.lang.Phonology)
		if err != nil {
			return nil, err
		}
		lwords[i] = *lw
	}

	return lwords, nil
}

func (fls *JSONFileLangStorage) DelWord(rom string) error {
	_, ind, err := fls.getWordIndex(rom)
	if err != nil {
		return err
	}
	fls.lang.Lexicon.Words = append(fls.lang.Lexicon.Words[:ind], fls.lang.Lexicon.Words[ind+1:]...)
	return nil
}

func (fls *JSONFileLangStorage) ParseWord(rom string) (*lang.Word, error) {
	lp := fls.lang.Phonology.toLang()
	return lp.ParseWord(rom)
}

func (fls *JSONFileLangStorage) GenerateWord(l string) (string, error) {
	lp := fls.lang.Phonology.toLang()
	return lp.GenerateWord(l)
}
