package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

/*
{"derivation": "语出《法华经·法师功德品》下至阿鼻地狱。”",
 "example": "但也有少数意志薄弱的……逐步上当，终至堕入～。★《上饶集中营·炼狱杂记》",
 "explanation": "阿鼻梵语的译音，意译为无间”，即痛苦无有间断之意。常用来比喻黑暗的社会和严酷的牢狱。又比喻无法摆脱的极其痛苦的境地。",
  "pinyin": "ā bí dì yù",
   "word": "阿鼻地狱",
   "abbreviation": "abdy"}
*/
type IdiomItem struct {
	Derivation   string
	Example      string
	Explanation  string
	Pinyin       string
	Word         string
	Abbreviation string
}

type GraphItem struct {
	Word string
	ID   int
	Next []int
}

func MakeGraph() []GraphItem {
	var res []IdiomItem
	raw, err := ioutil.ReadFile("./files/idiom.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(raw, &res)
	if err != nil {
		log.Fatal(err)
	}
	var graph = make([]GraphItem, len(res))
	for i := 0; i < len(res); i++ {
		graph[i] = GraphItem{Word: res[i].Word, ID: i}
	}
	for i := 0; i < len(res); i++ {
		var next []int
		word1 := []rune(res[i].Word)
		for j := 0; j < len(res); j++ {
			word2 := []rune(res[j].Word)
			if word2[0] == word1[len(word1)-1] && i != j {
				next = append(next, graph[j].ID)
			}
		}
		graph[i].Next = next
	}
	return graph
}

func WriteGraph(graph []GraphItem) {
	data, err := json.Marshal(graph)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("./files/graph.json", data, 0644)
}
func main() {
	var graph []GraphItem = MakeGraph()
	WriteGraph(graph)
}
