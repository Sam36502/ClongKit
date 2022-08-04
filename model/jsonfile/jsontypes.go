package jsonfile

type Language struct {
	Phonology Phonology `json:"phonology"`
	Lexicon   Lexicon   `json:"lexicon"`
}

type Phonology struct {
	Phonemes []Phoneme `json:"phonemes"`
}

type Phoneme struct {
	IPA          string `json:"ipa"`
	Romanisation string `json:"rom"`
	GroupID      string `json:"grp"`
}

type Lexicon struct {
	Words []Word `json:"words"`
}

type Word struct {
	Romanisation string   `json:"rom"`
	Etymology    string   `json:"ety"`
	Meanings     []string `json:"mns"`
	Tags         []string `json:"tgs"`
}
