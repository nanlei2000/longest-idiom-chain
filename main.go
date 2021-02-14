package main

import (
	"fmt"
	"os"

	"log"

	f "github.com/nanlei2000/longest-idiom-chain/findchain"
	"github.com/urfave/cli/v2"
)

func main() {

	var maxLoopCount int64
	app := &cli.App{
		Name:  "findchain",
		Usage: "find the longest idiom chain",
		Flags: []cli.Flag{
			&cli.Int64Flag{
				Name:        "mlc",
				Value:       100_000,
				Usage:       "max loop count for dfs",
				Destination: &maxLoopCount,
			},
		},
		Action: func(c *cli.Context) error {
			idiom := "文武双全"
			if c.NArg() > 0 {
				idiom = c.Args().Get(0)
			}
			fmt.Printf("[info] head idiom: %s, max loop count: %v\n", idiom, maxLoopCount)
			graph := f.ReadGraph()
			idToGraphItemMap := f.MakeIDToGraphItemMap(graph)
			wordToGraphItemMap := f.MakeWordToGraphItemMap(graph)
			wordID := wordToGraphItemMap[idiom].ID
			longest := f.FindLongestChain(wordID, idToGraphItemMap, maxLoopCount)
			words := f.MapIDtoIdiom(longest, idToGraphItemMap)
			fmt.Printf("chain: %s\n", words)
			fmt.Printf("length: %v\n", len(longest))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
