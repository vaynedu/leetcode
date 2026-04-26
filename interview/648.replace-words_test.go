package interview

import "testing"

func TestReplaceWords(t *testing.T) {
	dictionary := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the baby"

	result := ReplaceWords(dictionary, sentence)
	expected := "the cat was rat by the baby"

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestReplaceWordsNoRoot(t *testing.T) {
	dictionary := []string{"a"}
	sentence := "a b c d"

	result := ReplaceWords(dictionary, sentence)
	expected := "a b c d"

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestReplaceWordsEmpty(t *testing.T) {
	dictionary := []string{}
	sentence := "hello world"

	result := ReplaceWords(dictionary, sentence)
	if result != sentence {
		t.Errorf("expected %q, got %q", sentence, result)
	}
}

func TestReplaceWordsExactMatch(t *testing.T) {
	dictionary := []string{"cat"}
	sentence := "cat"

	result := ReplaceWords(dictionary, sentence)
	if result != "cat" {
		t.Errorf("expected 'cat', got %q", result)
	}
}
