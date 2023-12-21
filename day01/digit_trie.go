package main

import "unicode"

type DigitNode struct {
	children [26]*DigitNode
	isWord   bool
}

type DigitTrie struct {
	root *DigitNode
}

func InitDigitTrie() *DigitTrie {
	result := &DigitTrie{root: &DigitNode{}}
	result.insert("one")
	result.insert("two")
	result.insert("three")
	result.insert("four")
	result.insert("five")
	result.insert("six")
	result.insert("seven")
	result.insert("eight")
	result.insert("nine")
	return result
}

func (t *DigitTrie) insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &DigitNode{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isWord = true
}

/*
returns pointer to the next node or null
*/
func (t *DigitTrie) IsStartOfDigit(r rune) *DigitNode {
	currentNode := t.root
	charIndex := r - 'a'
	if currentNode.children[charIndex] == nil {
		return nil
	}
	nextNode := currentNode.children[charIndex]
	return nextNode
}

func (t *DigitTrie) IsNextCharInDigit(node *DigitNode, r rune) *DigitNode {
	charIndex := r - 'a'
	if unicode.IsNumber(r) || node.children[charIndex] == nil {
		return nil
	}
	nextNode := node.children[charIndex]
	return nextNode
}
