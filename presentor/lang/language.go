package lang

type Language struct {
	Phonology Phonology `json:"phono"`
	Lexicon   Lexicon   `json:"lex"`
}
