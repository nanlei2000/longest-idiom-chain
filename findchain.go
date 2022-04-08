package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
)

// GraphItem 图的节点
type GraphItem struct {
	Word string
	ID   uint16
	Next []uint16
}

// TNodeMap define id -> GraphItem map type
type TNodeMap map[uint16]GraphItem

// TWordToGraphItemMap define word -> GraphItem map type
type TWordToGraphItemMap map[string]GraphItem

// ReadGraph read json file into go graph struct
func ReadGraph() []GraphItem {
	var res []GraphItem
	err := json.Unmarshal(Db, &res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

// MakeIDToGraphItemMap make a id->Graph map
func MakeIDToGraphItemMap(graph []GraphItem) TNodeMap {
	nodeMap := make(TNodeMap)
	for _, v := range graph {
		nodeMap[v.ID] = v
	}
	return nodeMap
}

// MakeWordToGraphItemMap make a word->Graph map
func MakeWordToGraphItemMap(graph []GraphItem) TWordToGraphItemMap {
	nodeMap := make(TWordToGraphItemMap)
	for _, v := range graph {
		nodeMap[v.Word] = v
	}
	return nodeMap
}

// FindLongestChain perform a dfs into Graph to find longet idiom chain
func FindLongestChain(id uint16, nodeMap TNodeMap, maxLoopCount int64) []uint16 {
	var dfs func(uint16, []uint16) []uint16

	dfs = func(id uint16, chain []uint16) []uint16 {
		maxLoopCount--
		if maxLoopCount < 0 {
			return chain
		}

		next := nodeMap[id].Next
		validNextWords := []uint16{}

		for _, v := range next {
			var isContain bool
			for _, vc := range chain {
				if vc == v {
					isContain = true
					break
				}
			}
			if !isContain {
				validNextWords = append(validNextWords, v)
			}
		}
		if len(validNextWords) == 0 {
			return chain
		}

		var maxLength = -1
		longestChain := []uint16{}

		for _, id := range validNextWords {
			path := []uint16{}
			path = append(path, chain...)
			path = append(path, id)
			currentChain := dfs(id, path)
			currentLength := len(currentChain)
			if currentLength > maxLength {
				maxLength = currentLength
				longestChain = currentChain
			}
		}
		return longestChain
	}
	return dfs(id, []uint16{id})
}

// MapIDtoIdiom map id list back string list
func MapIDtoIdiom(chain []uint16, nodeMap TNodeMap) []string {
	res := []string{}
	for _, id := range chain {
		res = append(res, nodeMap[id].Word)
	}
	return res
}

// WriteWordsFile write a json file to disk for keeping result
func WriteWordsFile(words []string) {
	filename := "./files/" +
		words[0] +
		"-" +
		strconv.FormatInt(int64(len(words)), 10) +
		".json"
	data, err := json.Marshal(words)
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		log.Fatal(err)
	}
}
