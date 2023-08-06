package pronouns

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsePronounString(t *testing.T) {
	p, err := ParsePronounString("he/him/his/his/himself")
	assert.NoError(t, err)
	assert.Equal(t, HeHim, p)
	p, err = ParsePronounString("she/her/her/hers/herself")
	assert.NoError(t, err)
	assert.Equal(t, SheHer, p)
	p, err = ParsePronounString("a/b/c/d/e")
	assert.NoError(t, err)
	assert.Equal(t, Pronoun{"a", "b", "c", "d", "e"}, p)
	_, err = ParsePronounString("he/him/his/his/")
	assert.Error(t, err)
	_, err = ParsePronounString("/him/his/his/himself")
	assert.Error(t, err)
	_, err = ParsePronounString("/him/his/his")
	assert.Error(t, err)
}

func TestPronoun_String(t *testing.T) {
	assert.Equal(t, "he/him", HeHim.String())
	assert.Equal(t, "she/her", SheHer.String())
	assert.Equal(t, "they/them", TheyThem.String())
	assert.Equal(t, "it/its", ItIts.String())
	assert.Equal(t, "one/one's", OneOnes.String())
	assert.Equal(t, "a/b", Pronoun{Subject: "a", Object: "b"}.String())
	assert.Equal(t, "a/b", Pronoun{Subject: "a", Object: "a", Possessive: "b"}.String())
}

func TestPronoun_Long(t *testing.T) {
	assert.Equal(t, "he/him/his/his/himself", HeHim.Long())
	assert.Equal(t, "she/her/her/hers/herself", SheHer.Long())
	assert.Equal(t, "a/b/c/d/e", Pronoun{"a", "b", "c", "d", "e"}.Long())
}
