package pkg

type Node struct {
	Term     bool
	Value    string
	Children map[string]*Node
}

type Trie struct {
	Root *Node
}

func New() *Trie {
	return &Trie{Root: &Node{}}
}
