package conllu

import (
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	type test struct {
		input     string
		expected  Token
		isComment bool
		isSep     bool
		shouldErr bool
	}

	tests := []test{
		{
			input: "11	the	the	DET	DT	Definite=Def|PronType=Art	12	det	_	Entity=(abstract-91",
			expected: Token{
				ID:    11,
				Form:  "the",
				Lemma: "the",
				UPOS:  "DET",
				XPOS:  "DT",
				Feats: []MorphologicalFeature{
					{
						Feature: "Definite",
						Value:   "Def",
					},
					{
						Feature: "PronType",
						Value:   "Art",
					},
				},
				Head:   12,
				Deprel: "det",
				Deps:   nil,
				Misc: []string{
					"Entity=(abstract-91",
				},
			},
		},
		{
			input:     "not a valid line",
			shouldErr: true,
		},
		{
			input:    "# a comment",
			isSep:    false,
			expected: Token{},
		},
		{
			input: "",
			isSep: true,
		},
	}

	for _, test := range tests {
		actual, isComment, isSep, err := parseLine(test.input)
		if test.shouldErr && err != nil {
			continue
		}
		if err != nil {
			t.Error(err)
		}
		if test.isSep {
			if !isSep {
				t.Errorf("expected sep: %v, actual: %v", test.isSep, isSep)
			}
			continue
		}
		if test.isComment {
			if !isComment {
				t.Errorf("expected comment: %v, actual: %v", test.isComment, isComment)
			}
			continue
		}
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("expected: %v, actual: %v", test.expected, actual)
		}
	}
}

func TestParseFeats(t *testing.T) {
	type test struct {
		input     string
		expected  []MorphologicalFeature
		shouldErr bool
	}

	tests := []test{
		{
			input: "Number=Sing",
			expected: []MorphologicalFeature{
				{
					Feature: "Number",
					Value:   "Sing",
				},
			},
		},
		{
			input: "Definite=Def|PronType=Art",
			expected: []MorphologicalFeature{
				{
					Feature: "Definite",
					Value:   "Def",
				},
				{
					Feature: "PronType",
					Value:   "Art",
				},
			},
		},
		{
			input:    "_",
			expected: nil,
		},
		{
			input:     "",
			shouldErr: true,
		},
	}

	for _, test := range tests {
		actual, err := parseFeats(test.input)
		if test.shouldErr {
			if err == nil {
				t.Errorf("should have errored, input: %v", test.input)
			}
			continue
		}
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("expected: %v, actual: %v", test.expected, actual)
		}
	}
}

func TestParseDeps(t *testing.T) {
	type test struct {
		input     string
		expected  []Dep
		shouldErr bool
	}

	tests := []test{
		{
			input: "0:root",
			expected: []Dep{
				{
					Head:   0,
					Deprel: "root",
				},
			},
		},
		{
			input: "2:nsubj|4:nsubj",
			expected: []Dep{
				{
					Head:   2,
					Deprel: "nsubj",
				},
				{
					Head:   4,
					Deprel: "nsubj",
				},
			},
		},
		{
			input:    "_",
			expected: nil,
		},
		{
			input:     "",
			shouldErr: true,
		},
	}

	for _, test := range tests {
		actual, err := parseDeps(test.input)
		if test.shouldErr {
			if err == nil {
				t.Errorf("should have errored, input: %v", test.input)
			}
			continue
		}
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("expected: %v, actual: %v", test.expected, actual)
		}
	}
}

func TestParseMisc(t *testing.T) {
	type test struct {
		input     string
		expected  []string
		shouldErr bool
	}

	tests := []test{
		{
			input: "thingy",
			expected: []string{
				"thingy",
			},
		},
		{
			input: "Entity=(abstract-96)|SpaceAfter=No",
			expected: []string{
				"Entity=(abstract-96)",
				"SpaceAfter=No",
			},
		},
		{
			input:    "_",
			expected: nil,
		},
		{
			input:     "",
			shouldErr: true,
		},
	}

	for _, test := range tests {
		actual, err := parseMisc(test.input)
		if test.shouldErr {
			if err == nil {
				t.Errorf("should have errored, input: %v", test.input)
			}
			continue
		}
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("expected: %v, actual: %v", test.expected, actual)
		}
	}
}
