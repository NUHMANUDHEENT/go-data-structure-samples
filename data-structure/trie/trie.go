package main

import "fmt"

type TrieNode struct {
	Children map[rune]*TrieNode
	Isend    bool
}
type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			Children: map[rune]*TrieNode{}}}
}
func (t *Trie) Insert(word string) {
	node := t.root
	for _, v := range word {
		if _, exist := node.Children[v]; !exist {
			node.Children[v] = &TrieNode{Children: make(map[rune]*TrieNode)}
		}
		node = node.Children[v]
	}
	node.Isend = true
}
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, v := range word {
		if _, exist := node.Children[v]; !exist {
			return false
		}
		node = node.Children[v]
	}
	return node.Isend
}
func main() {

	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("world")

	fmt.Println("Search for 'hello':", trie.Search("hello"))   // true
	fmt.Println("Search for 'hell':", trie.Search("hell"))     // false
	// fmt.Println("StartsWith 'hell':", trie.StartsWith("hell")) // true
	// fmt.Println("StartsWith 'wor':", trie.StartsWith("wor"))   // true
	// fmt.Println("StartsWith 'word':", trie.StartsWith("word")) // false
}
