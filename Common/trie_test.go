package Common_test

import (
	"github.com/GerardoHP/AdventOfCode2023/Common"
	"testing"
)

func TestNewTrieNode(t *testing.T) {
	trie := Common.NewTrieNode()
	trie.Insert("four")
	trie.Insert("five")
	trie.Insert("1")

	r1, v1 := trie.Search("four")
	r2, v2 := trie.Search("five")
	r3, v3 := trie.Search("foura")
	r4, v4 := trie.Search("1")

	if !r1 && v1 != 4 {
		t.Fatalf("r1 expected to be found")
	}

	if !r2 && v2 != 5 {
		t.Fatalf("r2 expected to be found")
	}

	if r3 && v3 != -1 {
		t.Fatalf("r3 expected to not be found")
	}

	if !r4 && v4 != 1 {
		t.Fatalf("r4 expecte te be found")
	}
}
