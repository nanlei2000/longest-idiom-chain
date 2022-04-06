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
	ID   int
	Next []int
}

// TNodeMap define id -> GraphItem map type
type TNodeMap map[int]GraphItem

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
func FindLongestChain(id int, nodeMap TNodeMap, maxLoopCount int64) []int {
	var dfs func(int, []int) []int
	// maxLoopCount := 100_000
	dfs = func(id int, chain []int) []int {
		maxLoopCount--
		if maxLoopCount < 0 {
			return chain
		}

		nextWordIDList := nodeMap[id].Next
		validNextWords := []int{}
		for _, v := range nextWordIDList {
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
		longestChain := []int{}
		for _, id := range validNextWords {
			path := []int{}
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
	return dfs(id, []int{id})
}

// MapIDtoIdiom map id list back string list
func MapIDtoIdiom(chain []int, nodeMap TNodeMap) []string {
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
