package trie

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{Root: NewTrieNode()}
}

func (t *Trie) Insert(word string) {
	current_node := t.Root

	for _, char := range word {
		s := string(char)
		if _, exists := current_node.Children[s]; !exists {
			current_node.Children[s] = NewTrieNode()
		}

		current_node = current_node.Children[s]
	}

	current_node.Is_End_Of_Word = true

}

func (t *Trie) Search(word string) bool {
	current_node := t.Root

	for _, char := range word {
		s := string(char)

		if _, exists := current_node.Children[s]; !exists {
			return false
		}
		current_node = current_node.Children[s]

	}

	return current_node.Is_End_Of_Word

}

func (t *Trie) ShowAll() []string {
	if t.Root == nil {
		return []string{}
	}
	// We call the node's method, starting with an empty prefix
	return t.Root.findWords("")
}

func (t *Trie) Starts_with(prefix string) bool {
	current_node := t.Root

	for _, char := range prefix {
		s := string(char)

		if _, exists := current_node.Children[s]; !exists {
			return false
		}
		current_node = current_node.Children[s]
	}

	return true
}
