package conllu

// Token represents a single token, e.g. "hello", "goodby"
// and holds all associated annotations
// https://universaldependencies.org/format.html#conll-u-format
type Token struct {
	ID uint // Word index, integer starting at 1 for each new sentence; may be a range for multiword tokens; may be a decimal number for empty nodes (decimal numbers can be lower than 1 but must be greater than 0)

	Form string // Word form or punctuation symbol

	Lemma string // Lemma or stem of word form

	UPOS string // Universal part-of-speech tag

	XPOS *string // Language-specific part-of-speech tag; nil if not available

	// List of morphological features, which are described on the type; nil if not available
	Feats []MorphologicalFeature

	// Head of the current word, which is either the id of the head token for this word, or nil if none
	// https://universaldependencies.org/format.html#syntactic-annotation
	Head uint

	// Universal dependency relation to the HEAD (root iff HEAD = 0) or a defined language-specific subtype of one
	Deprel string

	// Enhanced dependency graph in the form of a list of head-deprel pairs. See Dep type for more information; nil if none.
	// Dependencies that are shared between the basic and the enhanced dependency representations must be repeated in the Deps field
	Deps []Dep

	// Any other annotation, represented as a list separated by "|". Nil if none.
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
	Head   uint
	Deprel string
}