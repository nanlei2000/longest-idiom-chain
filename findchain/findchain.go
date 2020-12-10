package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// GraphItem 图的节点
type GraphItem struct {
	Word string
	ID   int
	Next []int
}

type tNodeMap map[int]GraphItem

func readGraph() []GraphItem {
	var res []GraphItem
	raw, err := ioutil.ReadFile("./files/graph.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(raw, &res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func makeMap(graph []GraphItem) tNodeMap {
	nodeMap := make(tNodeMap)
	for _, v := range graph {
		nodeMap[v.ID] = v
	}
	return nodeMap
}

func findLongestChain(id int, nodeMap tNodeMap) []int {
	var dfs func(int, []int) []int
	maxLoopCount := 1000000
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
			for _, v := range chain {
				path = append(path, v)
			}
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
	return dfs(id, []int{})
}
func main() {
	graph := readGraph()
	nodeMap := makeMap(graph)
	longest := findLongestChain(22989, nodeMap)
	fmt.Println(len(longest))
}
