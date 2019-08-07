package algorithms

import "strings"

// TrieNode node type
type TrieNode struct {
	char        rune
	isWholeWord bool
	children    map[rune]*TrieNode
}

// Trie implementation
type Trie struct {
	root *TrieNode
}

// Insert a word into the trie
func (trie *Trie) Insert(word string) {
	if trie.root == nil {
		root := TrieNode{}
		root.children = make(map[rune]*TrieNode)
		trie.root = &root
	}

	crawl := trie.root

	for _, char := range word {
		_, ok := crawl.children[char]
		if !ok {
			node := TrieNode{char: char}
			node.children = make(map[rune]*TrieNode)
			crawl.children[char] = &node

		}
		crawl = crawl.children[char]
	}

	crawl.isWholeWord = true
}

// Contains checks if the word is in the Trie
func (trie *Trie) Contains(word string) bool {
	node := trie.findNode(word)
	return node != nil && node.isWholeWord
}

// StartsWith returns if there is any word in the trie
// that starts with the given prefix.
func (trie *Trie) StartsWith(prefix string) bool {
	node := trie.findNode(prefix)
	return node != nil
}

func (trie *Trie) findNode(word string) *TrieNode {
	crawl := trie.root
	for _, char := range word {
		node, ok := crawl.children[char]
		if !ok {
			return nil
		}
		crawl = node
	}
	return crawl
}

// LongestCommonPrefix : https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/887/
func (trie *Trie) LongestCommonPrefix() string {
	crawl := trie.root
	var sb strings.Builder
	for len(crawl.children) == 1 && !crawl.isWholeWord {
		for char, node := range crawl.children {
			crawl = node
			sb.WriteRune(char)
		}
	}
	return sb.String()
}
