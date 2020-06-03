package conllu

import (
	"bufio"
	"io"
	"os"
)

// Parse parses conllu via the io.Reader interface and returns all of the tokens found
// Parse doesn't close the reader when finished, that must be done manually
func Parse(r io.Reader) ([]Sentence, error) {
	line := ""
	lineNumber := 0
	var err error
	sentences := []Sentence{}
	currentSentence := Sentence{
		Tokens: []Token{},
	}
	reader := bufio.NewReader(r)
	for err == nil {
		lineNumber++
		line, err = reader.ReadString('\n')
		token, isComment, isSep, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		if isComment {
			continue
		}
		if isSep {
			if len(currentSentence.Tokens) == 0 {
				continue
			}
			sentences = append(sentences, currentSentence)
			currentSentence = Sentence{
				Tokens: []Token{},
			}
			continue
		}
		currentSentence.Tokens = append(currentSentence.Tokens, token)
	}
	return sentences, nil
}

// ParseFile opens, reads, and parses a file in conllu format and returns all of the tokens found.
// ParseFile is a convencience wrapper for the Parse() function when working with files on disk
func ParseFile(filepath string) ([]Sentence, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return Parse(file)
}
