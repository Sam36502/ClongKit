package model

import "github.com/Sam36502/ClongKit/lang"

type LangStorage interface {

	// Language Methods
	CreateLanguage(name, ID string) error
	GetName() (string, error)
	GetID() (string, error)
	SetName(string) error
	SetID(string) error
	Close() error

	// Phonology Methods
	SetPhoneme(ph lang.Phoneme) error
	GetPhoneme(rom string) (*lang.Phoneme, error)
	GetAllPhonemes() ([]lang.Phoneme, error)
	GetPhonology() (lang.Phonology, error)
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
}
