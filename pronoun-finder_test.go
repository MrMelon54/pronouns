package pronouns

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPronounFinder(t *testing.T) {
	assert.Equal(t, map[string]Pronoun{
		"he/him":    HeHim,
		"it/its":    ItIts,
		"one/one's": OneOnes,
		"she/her":   SheHer,
		"they/them": TheyThem,
	}, defaultFinder.pronouns)
	assert.Equal(t, map[string]Pronoun{
		"they/them": TheyThem,
	}, NewPronounFinder([]Pronoun{TheyThem}, false).pronouns)
}

func TestFindPronoun(t *testing.T) {
	p, err := FindPronoun("he/him")
	assert.NoError(t, err)
	assert.Equal(t, HeHim, p)
	p, err = FindPronoun("she/her")
	assert.NoError(t, err)
	assert.Equal(t, SheHer, p)
	p, err = FindPronoun("they/them")
	assert.NoError(t, err)
	assert.Equal(t, TheyThem, p)
	p, err = FindPronoun("it/its")
	assert.NoError(t, err)
	assert.Equal(t, ItIts, p)
	p, err = FindPronoun("one/one's")
	assert.NoError(t, err)
	assert.Equal(t, OneOnes, p)
}
