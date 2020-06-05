package conllu

// Sentence represents a single sentence of parsed CoNLL-U tokens
type Sentence struct {
	Tokens []Token
}

// Token represents a single token, e.g. "hello", "bye"
// and holds all associated annotations
// https://universaldependencies.org/format.html#conll-u-format
type Token struct {
	// Word index, float starting at 1 for each new sentence
	// If a range was found on a single line in the file
	// then nultiple tokens will be created
	ID float64

	// Word form or punctuation symbol
	Form string

	// Lemma or stem of word form
	Lemma string

	// Universal part-of-speech tag
	UPOS string

	// Language-specific part-of-speech tag; empty if not available
	XPOS string

	// List of morphological features, which are described on the type
	// nil if not available
	Feats []MorphologicalFeature

	// Head of the current word
	// Either the ID of the head token for this word, or 0 if root
	// https://universaldependencies.org/format.html#syntactic-annotation
	Head float64

	// Universal dependency relation to the HEAD (root iff HEAD = 0)
	// or a defined language-specific subtype of one
	Deprel string

	// Enhanced dependency graph in the form of a list of head-deprel pairs. See Dep type for more information; nil if none.
	// Dependencies that are shared between the basic and the enhanced dependency representations must be repeated in the Deps field
	Deps []Dep

	// Any other annotation, nil if none.
	// https://universaldependencies.org/format.html#miscellaneous
	Misc []string
}

// MorphologicalFeature from the universal feature inventory (https://universaldependencies.org/u/feat/index.html)
// or from a defined language-specific extension (https://universaldependencies.org/ext-feat-index.html)
type MorphologicalFeature struct {
	Feature string
	Value   string
}

// Dep is a representation of a single part of the enhanced dependency graph
type Dep struct {
	Head   float64
	Deprel string
}
