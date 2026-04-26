package interview

import (
	"slices"
	"testing"
)

func TestFindWords212(t *testing.T) {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}

	result := FindWords(board, words)
	slices.Sort(result)

	expected := []string{"eat", "oath"}
	if len(result) != len(expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("at index %d: expected %s, got %s", i, expected[i], result[i])
		}
	}
}

func TestFindWords212Single(t *testing.T) {
	board := [][]byte{{'a', 'a'}}
	words := []string{"aaa"}

	result := FindWords(board, words)
	if len(result) != 0 {
		t.Errorf("expected empty, got %v", result)
	}
}

func TestFindWords212AllFound(t *testing.T) {
	board := [][]byte{
		{'a'},
	}
	words := []string{"a"}

	result := FindWords(board, words)
	if len(result) != 1 || result[0] != "a" {
		t.Errorf("expected [a], got %v", result)
	}
}
