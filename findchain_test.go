package main

import (
	"testing"
)

func TestFindLongestChain(t *testing.T) {
	graph := ReadGraph()
	nodeMap := MakeIDToGraphItemMap(graph)
	longest := FindLongestChain(23192, nodeMap, 100000)
	words := MapIDtoIdiom(longest, nodeMap)

	t.Logf("words: %s", words)

	isValidChain := true
	for i, v := range words {
		if i == 0 {
			continue
		} else {
			prevWordRune := []rune(words[i-1])
			currentWordRune := []rune(v)
			if currentWordRune[0] != prevWordRune[len(prevWordRune)-1] {
				isValidChain = false
				break
			}
		}
	}

	if !isValidChain {
		t.Fatalf("invalid chain: %s", words)
	}
}
