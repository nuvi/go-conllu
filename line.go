package conllu

import (
	"fmt"
	"strconv"
	"strings"
)

// return (t Token, isComment, isSep bool, err error)
func parseLine(line string) (Token, bool, bool, error) {
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return Token{}, false, true, nil
	}

	if string(line[0]) == "#" {
		return Token{}, true, false, nil
	}

	entries := strings.Split(line, "\t")
	if len(entries) != 10 {
		return Token{}, false, false, fmt.Errorf("line %v has incorrect number of entries. expected 10 found %v", entries, len(entries))
	}

	t := Token{}
	id, err := strconv.ParseFloat(entries[0], 64)
	if err != nil {
		return Token{}, false, false, fmt.Errorf("ID can't be parsed. id: %v, err: %v", entries[0], err)
	}
	t.ID = id

	form := entries[1]
	if form == "" {
		return Token{}, false, false, fmt.Errorf("FORM can't be parsed. form: %v, err: %v", form, err)
	}
	t.Form = form

	lemma := entries[2]
	if lemma == "" {
		return Token{}, false, false, fmt.Errorf("LEMMA can't be parsed. lemma: %v, err: %v", lemma, err)
	}
	t.Lemma = lemma

	upos := entries[3]
	if upos == "" {
		return Token{}, false, false, fmt.Errorf("UPOS can't be parsed. upos: %v, err: %v", upos, err)
	}
	t.UPOS = upos

	xpos := entries[4]
	if xpos == "" {
		return Token{}, false, false, fmt.Errorf("UPOS can't be parsed. xpos: %v, err: %v", xpos, err)
	}
	if xpos != "_" {
		t.XPOS = xpos
	}

	feats := entries[5]
	finalFeats, err := parseFeats(feats)
	if err != nil {
		return Token{}, false, false, err
	}
	t.Feats = finalFeats

	if entries[6] == "_" {
		t.Head = 0
	} else {
		head, err := strconv.ParseFloat(entries[6], 64)
		if err != nil {
			return Token{}, false, false, fmt.Errorf("HEAD can't be parsed. id: %v, err: %v", entries[6], err)
		}
		t.Head = head
	}

	deprel := entries[7]
	if deprel == "_" {
		deprel = "root"
	}
	if t.Head == 0 && deprel != "root" {
		return Token{}, false, false, fmt.Errorf("DEPREL must match head. deprel: %v, head: %v", deprel, t.Head)
	}
	t.Deprel = deprel

	deps := entries[8]
	finalDeps, err := parseDeps(deps)
	if err != nil {
		return Token{}, false, false, err
	}
	t.Deps = finalDeps

	misc := entries[9]
	finalMisc, err := parseMisc(misc)
	if err != nil {
		return Token{}, false, false, err
	}
	t.Misc = finalMisc
	return t, false, false, nil
}

func parseFeats(feats string) ([]MorphologicalFeature, error) {
	if feats == "_" {
		return nil, nil
	}
	separated := strings.Split(feats, "|")
	finalFeatures := []MorphologicalFeature{}
	for _, sep := range separated {
		pieces := strings.Split(sep, "=")
		if len(pieces) != 2 {
			return nil, fmt.Errorf("Invalid FEAT length. text: %v, len: %v", sep, len(pieces))
		}
		finalFeatures = append(finalFeatures, MorphologicalFeature{
			Feature: pieces[0],
			Value:   pieces[1],
		})
	}
	return finalFeatures, nil
}

func parseDeps(deps string) ([]Dep, error) {
	if deps == "_" {
		return nil, nil
	}
	separated := strings.Split(deps, "|")
	finalDeps := []Dep{}
	for _, sep := range separated {
		pieces := strings.Split(sep, ":")
		if len(pieces) < 2 {
			return nil, fmt.Errorf("Invalid DEP length. text: %v, len: %v", sep, len(pieces))
		}
		head, err := strconv.ParseFloat(pieces[0], 64)
		if err != nil {
			return nil, fmt.Errorf("DEPS HEAD can't be parsed. id: %v, err: %v", pieces[0], err)
		}
		finalDeps = append(finalDeps, Dep{
			Head:   head,
			Deprel: pieces[1],
		})
	}
	return finalDeps, nil
}

func parseMisc(values string) ([]string, error) {
	if values == "_" {
		return nil, nil
	}
	separated := strings.Split(values, "|")
	finalMiscs := []string{}
	for _, sep := range separated {
		if sep == "" {
			return nil, fmt.Errorf("Invalid MISC length. text: %v, len: %v", sep, len(sep))
		}
		finalMiscs = append(finalMiscs, sep)
	}
	return finalMiscs, nil
}
