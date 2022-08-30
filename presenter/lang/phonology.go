package lang

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Phoneme struct {
	IPA          string
	Romanisation string
	Groups       []string
}

type Phonology struct {
	Phonemes      []Phoneme
	SyllableRules []SyllableRule
}

func (ph *Phonology) ParseWord(rom string) (*Word, error) {
	newWord := Word{
		Phonemes: []Phoneme{},
	}
	for len(rom) > 0 {
		p, r, err := ph.parseOutFirstLetter(rom)
		if err != nil {
			return nil, err
		}
		rom = r
		newWord.Phonemes = append(newWord.Phonemes, *p)
	}
	return &newWord, nil
}

func (ph *Phonology) GetPhonemesByGroup(grp string) []Phoneme {
	phs := []Phoneme{}
	for _, p := range ph.Phonemes {
		for _, g := range p.Groups {
			if g == grp {
				phs = append(phs, p)
				break
			}
		}
	}
	return phs
}

func (l *Phonology) GenerateWord(lrange string) (string, error) {
	rand.Seed(time.Now().UnixNano())

	var leng int
	min, max, err := parseRangeStr(lrange)
	if err == nil && min < max {
		fmt.Println("Genning word")
		leng = rand.Intn(max-min) + min
	} else {
		num, err := strconv.ParseInt(lrange, 10, 32)
		if err == nil {
			leng = int(num)
		} else {
			return "", errors.New(fmt.Sprintf("Invalid length range string '%s'.", lrange))
		}
	}

	sb := strings.Builder{}
	for i := 0; i < leng; i++ {
		sb.WriteString(l.GenerateSyllable())
	}
	return sb.String(), nil
}

func (l *Phonology) GenerateSyllable() string {
	patt := l.SyllableRules[rand.Intn(len(l.SyllableRules))]
	sb := strings.Builder{}
	for _, g := range patt.OnsetGroups {
		phs := l.GetPhonemesByGroup(g)
		sb.WriteString(phs[rand.Intn(len(phs))].Romanisation)
	}
	phs := l.GetPhonemesByGroup(patt.NucleusGroup)
	sb.WriteString(phs[rand.Intn(len(phs))].Romanisation)
	for _, g := range patt.CodaGroups {
		phs := l.GetPhonemesByGroup(g)
		sb.WriteString(phs[rand.Intn(len(phs))].Romanisation)
	}
	return sb.String()
}

func parseRangeStr(l string) (int, int, error) {
	sepInd := strings.IndexRune(l, RangeSeparator)
	if sepInd < 0 {
		return -1, -1, errors.New(fmt.Sprintf("No range separator (%c) in range string", RangeSeparator))
	}
	min, err := strconv.ParseInt(strings.TrimSpace(l[:sepInd]), 10, 32)
	if err != nil {
		return -1, -1, err
	}
	max, err := strconv.ParseInt(strings.TrimSpace(l[sepInd+1:]), 10, 32)
	if err != nil {
		return -1, -1, err
	}
	return int(min), int(max), nil
}

// Takes a romanisation and parses out the first letter it finds
// Returns the phoneme, the shortened string and a possible error
func (ph *Phonology) parseOutFirstLetter(romanisation string) (*Phoneme, string, error) {

	for i := len(romanisation); i >= 0; i-- {
		p, err := ph.ParsePhoneme(romanisation[:i])
		if err == nil {
			return p, romanisation[i:], nil
		}
	}

	return nil, romanisation, errors.New(fmt.Sprintf("No letter found within word string '%s'", romanisation))
}

// Takes a romanisation letter/cluster and returns the phoneme
// associated with that letter
func (ph *Phonology) ParsePhoneme(letter string) (*Phoneme, error) {
	for _, p := range ph.Phonemes {
		if strings.Compare(strings.ToLower(p.Romanisation), strings.ToLower(letter)) == 0 {
			return &p, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Failed to parse letter '%s'", letter))
}
