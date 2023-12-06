package Common

type TrieNode struct {
	children *map[uint8]*TrieNode
	isEnd    bool
	val      int
}

func NewTrieNode() *TrieNode {
	m := make(map[uint8]*TrieNode)
	return &TrieNode{
		children: &m,
		isEnd:    false,
		val:      -1,
	}
}

func (tn *TrieNode) Insert(str string) {
	current := tn
	for level := 0; level < len(str); level++ {
		index := str[level] - 'a'
		if _, found := (*current.children)[index]; !found {
			(*current.children)[index] = NewTrieNode()
		}

		current = (*current.children)[index]
	}

	if found, n := stringToInt(str); found {
		current.val = n
	}

	current.isEnd = true
}

func (tn *TrieNode) Search(str string) (bool, int) {
	current := tn
	for level := 0; level < len(str); level++ {
		index := str[level] - 'a'

		if _, found := (*current.children)[index]; !found {
			return false, -1
		}

		current = (*current.children)[index]
	}

	return current.isEnd, current.val
}

func stringToInt(str string) (bool, int) {
	switch str {
	case "0", "zero":
		return true, 0
	case "1", "one":
		return true, 1
	case "2", "two":
		return true, 2
	case "3", "three":
		return true, 3
	case "4", "four":
		return true, 4
	case "5", "five":
		return true, 5
	case "6", "six":
		return true, 6
	case "7", "seven":
		return true, 7
	case "8", "eight":
		return true, 8
	case "9", "nine":
		return true, 9
	}

	return false, -1
}
