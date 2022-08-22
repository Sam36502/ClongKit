package model

import "github.com/Sam36502/ClongKit/presenter/lang"

type LangStorage interface {

	// Phonology Methods
	SetPhoneme(ph lang.Phoneme) error
	GetPhoneme(rom string) (*lang.Phoneme, error)
	GetAllPhonemes() ([]lang.Phoneme, error)
	DelPhoneme(rom string) error
	AddSyllableRule(patt lang.SyllableRule) error
	GetAllSyllableRules() ([]lang.SyllableRule, error)
	DelSyllableRule(patt lang.SyllableRule) error

	// Lexicon Methods
	SetWord(wrd lang.Word) error
	GetWord(rom string) (*lang.Word, error)
	GetAllWords() ([]lang.Word, error)
	SearchWord(rom, etym string, means, tags []string) ([]lang.Word, error)
	DelWord(rom string) error

	// Misc Methods
	ParseWord(rom string) (*lang.Word, error)
	Close() error
}
