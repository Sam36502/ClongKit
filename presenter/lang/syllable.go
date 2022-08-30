package lang

import (
	"errors"
	"fmt"
	"strings"
)

type SyllableRule struct {
	OnsetGroups  []string
	NucleusGroup string
	CodaGroups   []string
}

func (sr *SyllableRule) String() string {
	bits := append(sr.OnsetGroups, sr.NucleusGroup)
	bits = append(bits, sr.CodaGroups...)
	return strings.Join(bits, "")
}

func (sr *SyllableRule) Validate(phs []Phoneme) error {
	grps := map[string][]Phoneme{}
	for _, p := range phs {
		for _, g := range p.Groups {
			grps[g] = append(grps[g], p)
		}
	}
	for _, s := range sr.OnsetGroups {
		_, exists := grps[s]
		if !exists {
			return errors.New(fmt.Sprintf("Phoneme group '%s' not found", s))
		}
	}
	_, exists := grps[sr.NucleusGroup]
	if !exists {
		return errors.New(fmt.Sprintf("Phoneme group '%s' not found", sr.NucleusGroup))
	}
	for _, s := range sr.CodaGroups {
		_, exists := grps[s]
		if !exists {
			return errors.New(fmt.Sprintf("Phoneme group '%s' not found", s))
		}
	}
	return nil
}
