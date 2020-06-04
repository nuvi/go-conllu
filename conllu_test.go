package conllu

import "testing"

func TestParseFile(t *testing.T) {
	_, err := ParseFile("test_data/en_ewt-ud-train.conllu")
	if err != nil {
		t.Error(err)
	}
}

func TestParseFileSmall(t *testing.T) {
	sentences, err := ParseFile("test_data/en_ewt-ud-train.small.conllu")
	if err != nil {
		t.Error(err)
	}
	expected := 5
	if len(sentences) != expected {
		t.Errorf("Expecting len %v, got %v", expected, len(sentences))
	}

	expected = 29
	if len(sentences[0].Tokens) != expected {
		t.Errorf("Expecting len %v, got %v, data: %v", expected, len(sentences[0].Tokens), sentences[0].Tokens)
	}
	expected = 18
	if len(sentences[1].Tokens) != expected {
		t.Errorf("Expecting len %v, got %v, data: %v", expected, len(sentences[1].Tokens), sentences[1].Tokens)
	}
	expected = 17
	if len(sentences[2].Tokens) != expected {
		t.Errorf("Expecting len %v, got %v, data: %v", expected, len(sentences[2].Tokens), sentences[2].Tokens)
	}
	expected = 16
	if len(sentences[3].Tokens) != expected {
		t.Errorf("Expecting len %v, got %v, data: %v", expected, len(sentences[3].Tokens), sentences[3].Tokens)
	}
	expected = 16
	if len(sentences[4].Tokens) != expected {
		t.Errorf("Expecting len %v, got %v, data: %v", expected, len(sentences[4].Tokens), sentences[4].Tokens)
	}
}
