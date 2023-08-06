package pronouns

import (
	"errors"
	"strings"
)

// empty is used internally for returning errors
var empty = Pronoun{}

// Binary and normative-ish forms from https://en.pronouns.page/pronouns
var (
	HeHim    = Pronoun{"he", "him", "his", "his", "himself"}
	SheHer   = Pronoun{"she", "her", "her", "hers", "herself"}
	TheyThem = Pronoun{"they", "them", "their", "theirs", "themself"}
	ItIts    = Pronoun{"it", "it", "its", "its", "itself"}
	OneOnes  = Pronoun{"one", "one", "one's", "one's", "oneself"}
)

var ErrInvalidPronounString = errors.New("invalid pronoun string")

// Pronoun holds the 5 main forms of a pronoun in English
type Pronoun struct {
	Subject, Object, Possessive, PossessivePronoun, Reflexive string
}

// ParsePronounString splits the string a/b/c/d/e into the 5 separate components
// in the Pronoun struct.
//
// This trims any spaces around the 5 separate components.
//
// Errors are returned as (empty, ErrInvalidPronounString)
func ParsePronounString(v string) (Pronoun, error) {
	split := strings.Split(v, "/")
	if len(split) != 5 {
		return empty, ErrInvalidPronounString
	}
	for i := range split {
		split[i] = strings.TrimSpace(split[i])
		if split[i] == "" {
			return empty, ErrInvalidPronounString
		}
	}
	return Pronoun{split[0], split[1], split[2], split[3], split[4]}, nil
}

// String returns the short form output: he/him, she/her etc...
func (p Pronoun) String() string {
	if p.Subject == p.Object {
		return p.Subject + "/" + p.Possessive
	}
	return p.Subject + "/" + p.Object
}

// Long returns the full 5 components
func (p Pronoun) Long() string {
	return p.Subject + "/" + p.Object + "/" + p.Possessive + "/" + p.PossessivePronoun + "/" + p.Reflexive
}
