package model

import "github.com/Sam36502/ClongKit/presentor/lang"

type LangStorage interface {

	// Lexicon Methods
	SetWord(wrd lang.Word) error
	GetWord(rom string) (*lang.Word, error)
	GetAllWords() ([]lang.Word, error)
	SearchWord(rom, pron, etym, means, tags string) ([]lang.Word, error)
	DelWord(rom string) error

	// Misc Methods
	ParseWord(rom string) (*lang.Word, error)
	Close() error
}
