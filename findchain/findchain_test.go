package findchain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLongestChain(t *testing.T) {
	graph := ReadGraph()
	nodeMap := MakeIDToGraphItemMap(graph)
	longest := FindLongestChain(23192, nodeMap, 100000)
	words := MapIDtoIdiom(longest, nodeMap)

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

	assert.Equal(t, true, isValidChain, "invalid idiom chain")
}

func TestGetCurrentDir(t *testing.T) {
	dir := GetCurrentDir()
	t.Logf("dir is: %s", dir)
	assert.DirExists(t, dir)
}
