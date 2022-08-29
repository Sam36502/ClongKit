package jsonfile

import "github.com/Sam36502/ClongKit/presenter/lang"

func (ph *Phonology) toLang() lang.Phonology {
	lp := lang.Phonology{
		Phonemes: make([]lang.Phoneme, len(ph.Phonemes)),
	}
	for i, p := range ph.Phonemes {
		lp.Phonemes[i] = lang.Phoneme(p)
	}
	return lp
}

func (w *Word) toLang(ph *Phonology) (*lang.Word, error) {
	pl := ph.toLang()
	wrd, err := pl.ParseWord(w.Romanisation)
	if err != nil {
		return nil, err
	}

	wrd.Etymology = w.Etymology
	wrd.Meanings = w.Meanings
	wrd.Tags = w.Tags

	return wrd, nil
}

func (sr *SyllableRule) toLang() lang.SyllableRule {
	return lang.SyllableRule{
		OnsetGroups:  sr.OnsetGroups,
		NucleusGroup: sr.NucleusGroup,
		CodaGroups:   sr.CodaGroups,
	}
}
