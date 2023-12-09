package main

import (
	"fmt"

	trie "github.com/chaithanyaKS/dictionary/pkg/trie"
)

func main() {
	trie := trie.New()
	fmt.Println(trie)
	trie.Add("test")
	trie.Add("team")
	trie.Add("train")
	trie.Print()
}
