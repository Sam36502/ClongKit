package lang

import (
	"strings"
)

type Word struct {
	Phonemes  []Phoneme
	Stress    int
	Meanings  []string
	Etymology string
	Tags      []string
}

func (w *Word) GetRomanisation() string {
	rom := strings.Builder{}
	for _, p := range w.Phonemes {
		rom.WriteString(p.Romanisation)
	}
	return rom.String()
}

func (w *Word) GetPronunciation() string {
	pron := strings.Builder{}
	for _, p := range w.Phonemes {
		pron.WriteString(p.IPA)
	}
	return pron.String()
}
