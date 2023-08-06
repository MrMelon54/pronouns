package pronouns

import (
	"errors"
	"strings"
)

var (
	standardPronouns  = []Pronoun{HeHim, SheHer, TheyThem, ItIts, OneOnes}
	defaultFinder     = NewPronounFinder([]Pronoun{}, true)
	errInvalidPronoun = errors.New("invalid pronoun")
	errUnknownPronoun = errors.New("unknown pronoun")
)

// PronounFinder is used to find the pronoun struct based on the short form
// representation.
type PronounFinder struct {
	pronouns map[string]Pronoun
}

// FindPronoun is a utility function which uses the defaultFinder
func FindPronoun(v string) (Pronoun, error) {
	return defaultFinder.Find(v)
}

// NewPronounFinder creates a PronounFinder using the listed pronouns. Set std to
// true to load the standardPronouns as well.
func NewPronounFinder(p []Pronoun, std bool) *PronounFinder {
	pronouns := make(map[string]Pronoun)

	// add the standardPronouns
	if std {
		for _, i := range standardPronouns {
			pronouns[i.String()] = i
		}
	}

	// add custom pronouns
	for _, i := range p {
		pronouns[i.String()] = i
	}
	return &PronounFinder{pronouns}
}

// Find locates the pronoun matching the input short string.
// "he/him", "he / him" -> HeHim
// "she/her", "she / her" -> SheHer
func (p *PronounFinder) Find(v string) (Pronoun, error) {
	// find the first slash, no slashes is an error
	n := strings.IndexByte(v, '/')
	if n == -1 {
		return empty, errInvalidPronoun
	}

	// trim the spaces
	a := strings.TrimSpace(v[:n])
	b := strings.TrimSpace(v[n+1:])

	// there should be no more slashes after this
	if strings.IndexByte(b, '/') != -1 {
		return empty, errInvalidPronoun
	}

	// reconstruct the short form
	full := a + "/" + b

	// locate the pronoun in the map
	if pronoun, ok := p.pronouns[full]; ok {
		return pronoun, nil
	}
	return empty, errUnknownPronoun
}
