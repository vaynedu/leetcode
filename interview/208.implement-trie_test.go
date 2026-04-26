package interview

import "testing"

func TestTrie208(t *testing.T) {
	trie := NewTrie208()

	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Error("apple should exist")
	}
	if trie.Search("app") {
		t.Error("app should not exist as full word")
	}
	if !trie.StartsWith("app") {
		t.Error("app prefix should exist")
	}

	trie.Insert("app")
	if !trie.Search("app") {
		t.Error("app should exist after insert")
	}
}
