package algorithms

import (
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	trie := Trie{}
	trie.Insert("apple")
	trie.Insert("ape")
	trie.Insert("april")
	trie.Insert("doggie")
	if !trie.Contains("apple") {
		t.Error()
	}
	if trie.Contains("approve") {
		t.Error()
	}
	if !trie.Contains("doggie") {
		t.Error()
	}
	if !trie.StartsWith("app") {
		t.Error()
	}
	if trie.StartsWith("some") {
		t.Error()
	}
}

func TestTrie_LongestCommonPrefix(t *testing.T) {
	trie := Trie{}
	trie.Insert("flower")
	trie.Insert("flow")
	trie.Insert("flight")
	if trie.LongestCommonPrefix() != "fl" {
		t.Errorf("Longest common prefix: %s", trie.LongestCommonPrefix())
	}

	trie = Trie{}
	trie.Insert("dog")
	trie.Insert("racecar")
	trie.Insert("car")
	if trie.LongestCommonPrefix() != "" {
		t.Errorf("Longest common prefix: %s", trie.LongestCommonPrefix())
	}
}
