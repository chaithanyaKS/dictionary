package pkg

import "fmt"

type Trie struct {
	Term     bool
	Value    rune
	Children map[rune]*Trie
}

func New() *Trie {
	return new(Trie)
}

func (t *Trie) Add(input string) {
	node := t
	for _, str := range input {
		if node.Children == nil {
			node.Children = map[rune]*Trie{}
		}
		child := node.Children[str]
		if child == nil {
			child = &Trie{Value: str}
			node.Children[str] = child
		}
		node = child
	}
}

func (t *Trie) Print() {
	node := t
	for _, value := range node.Children {
		t.printNode(value)
	}
}

func (t *Trie) printNode(node *Trie) {
	if node == nil {
		return
	}
	fmt.Printf("%q -> ", node.Value)
	for _, value := range node.Children {
		t.printNode(value)
	}
	fmt.Println()
}
