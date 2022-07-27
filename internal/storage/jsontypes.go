package storage

import "github.com/Sam36502/ClongKit/internal/lang"

type JSONLanguage struct {
	Phonology JSONPhonology `json:"phonology"`
	Lexicon   JSONLexicon   `json:"lexicon"`
}

type JSONPhonology struct {
	Phonemes []JSONPhoneme `json:"phonemes"`
}

type JSONPhoneme struct {
	IPA          string `json:"ipa"`
	Romanisation string `json:"rom"`
	GroupID      string `json:"grp"`
}

type JSONLexicon struct {
	Words []JSONWord `json:"words"`
}

type JSONWord struct {
	Romanisation string   `json:"rom"`
	Meanings     []string `json:"mns"`
	Etymology    string   `json:"ety"`
	Tags         []string `json:"tgs"`
}

func (jl *JSONLanguage) ToLanguage() (*lang.Language, error) {
	l := lang.Language{
		Phonology: lang.Phonology{
			Phonemes: make([]lang.Phoneme, len(jl.Phonology.Phonemes)),
		},
		Lexicon: lang.Lexicon{
			Words: make([]lang.Word, len(jl.Lexicon.Words)),
		},
	}
	for i, p := range jl.Phonology.Phonemes {
		l.Phonology.Phonemes[i] = lang.Phoneme(p)
	}
	for i, w := range jl.Lexicon.Words {
		parsedWord, err := l.ParseWord(w.Romanisation)
		if err != nil {
			return nil, err
		}
		l.Lexicon.Words[i] = lang.Word{
			Phonemes:  parsedWord.Phonemes,
			Meanings:  w.Meanings,
			Etymology: w.Etymology,
			Tags:      w.Tags,
		}
	}
	return &l, nil
}
