package conllu

import "io"

// Parse parses conllu via the io.Reader interface and returns all of the tokens found
func Parse(reader io.Reader) ([]Token, error) {
	return nil, nil
}

// ParseString a string in conllu format and returns all of the tokens found
// If large strings are to be processed, it is recommended to use the Parse
// function instead
func ParseString(conllu string) ([]Token, error) {
	return nil, nil
}

// ParseFile opens, reads, and parses a file in conllu format and returns all of the tokens found
func ParseFile(filepath string) ([]Token, error) {
	return nil, nil
}
