package Days

import (
	"github.com/GerardoHP/AdventOfCode2023/Common"
)

var (
	t = Common.NewTrieNode()
)

func Day1(codes []string) int {
	prepTrie()
	result := 0
	for _, v := range codes {
		result += decode(v)
	}

	return result
}

func prepTrie() {
	t.Insert("one")
	t.Insert("two")
	t.Insert("three")
	t.Insert("four")
	t.Insert("five")
	t.Insert("six")
	t.Insert("seven")
	t.Insert("eight")
	t.Insert("nine")
	t.Insert("zero")
	t.Insert("1")
	t.Insert("2")
	t.Insert("3")
	t.Insert("4")
	t.Insert("5")
	t.Insert("6")
	t.Insert("7")
	t.Insert("8")
	t.Insert("9")
	t.Insert("0")
}

func decode(code string) int {
	a := make([]int, 0)
	start := 0
	end := 0

	for start < len(code) {
		subCode := code[start : end+1]
		if found, n := t.Search(subCode); found {
			a = append(a, n)
			start = len(code)
			end = len(code)
			break
		}

		end++
		if end == len(code) {
			end = start + 1
			start++
		}
	}

	for start > 0 {
		subCode := code[end-1 : start]
		if found, n := t.Search(subCode); found {
			if len(a) < 2 {
				a = append(a, n)
				start = 0
				end = 0
				break
			}
		}

		end--
		if end == 0 {
			end = start - 1
			start--
		}
	}

	if len(a) == 0 {
		return 0
	}

	if len(a) == 1 {
		return a[0]*10 + a[0]
	}

	return a[0]*10 + a[1]
}
