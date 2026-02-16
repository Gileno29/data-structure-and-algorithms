package trie

type TrieNode struct {
	Children       map[string]*TrieNode
	Is_End_Of_Word bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children:       make(map[string]*TrieNode),
		Is_End_Of_Word: false,
	}

}

func (n *TrieNode) findWords(prefix string) []string {
	var results []string

	if n.Is_End_Of_Word {
		results = append(results, prefix)
	}

	for char, child := range n.Children {
		// We pass the current prefix + the child's character down
		results = append(results, child.findWords(prefix+char)...)
	}

	return results
}
