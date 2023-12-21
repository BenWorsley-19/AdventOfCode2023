package utils

type Node struct {
	children [26]*Node
	isWord   bool
}

type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isWord = true
}

func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isWord == true {
		return true
	}

	return false
}

// func main() {
// 	testTrie := InitTrie()
// 	testTrie.Insert("one")
// 	testTrie.Insert("two")
// 	testTrie.Insert("three")
// 	fmt.Print(testTrie.Search("hello"))
// 	fmt.Print(testTrie.Search("one"))
// }
