package lang

type Phoneme struct {
	IPA          string `json:"ipa"`
	Romanisation string `json:"rom"`
	GroupID      string `json:"grp"`
}

type Phonology struct {
	Phonemes []Phoneme `json:"phonemes"`
}
