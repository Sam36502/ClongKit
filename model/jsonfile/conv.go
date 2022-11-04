package jsonfile

import "github.com/Sam36502/ClongKit/lang"

func (ph *Phonology) toLang() lang.Phonology {
	lp := lang.Phonology{
		Phonemes:      make([]lang.Phoneme, len(ph.Phonemes)),
		SyllableRules: make([]lang.SyllableRule, len(ph.SyllableRules)),
	}
	for i, p := range ph.Phonemes {
		lp.Phonemes[i] = lang.Phoneme(p)
	}
	for i, s := range ph.SyllableRules {
		lp.SyllableRules[i] = lang.SyllableRule(s)
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
