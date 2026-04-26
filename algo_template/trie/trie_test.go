package trie

import "testing"

func TestTrieBasic(t *testing.T) {
	trie := NewTrie()

	// Insert
	trie.Insert("apple")
	trie.Insert("apricot")

	// Search
	if !trie.Search("apple") {
		t.Error("apple should exist")
	}
	if trie.Search("app") {
		t.Error("app should not exist as full word (not inserted)")
	}
	if !trie.Search("apricot") {
		t.Error("apricot should exist")
	}

	// 再插入 app，然后它就存在了
	trie.Insert("app")
	if !trie.Search("app") {
		t.Error("app should exist after insert")
	}

	// StartsWith
	if !trie.StartsWith("app") {
		t.Error("app prefix should exist")
	}
	if !trie.StartsWith("apr") {
		t.Error("apr prefix should exist")
	}
	if trie.StartsWith("ban") {
		t.Error("ban prefix should not exist")
	}
}

func TestTrieEmpty(t *testing.T) {
	trie := NewTrie()

	if trie.Search("a") {
		t.Error("empty trie should not contain 'a'")
	}
	if trie.StartsWith("a") {
		t.Error("empty trie should not start with 'a'")
	}
}

func TestTrieSingleChar(t *testing.T) {
	trie := NewTrie()
	trie.Insert("a")

	if !trie.Search("a") {
		t.Error("'a' should exist")
	}
	if !trie.StartsWith("a") {
		t.Error("'a' prefix should exist")
	}
	if trie.Search("b") {
		t.Error("'b' should not exist")
	}
}
