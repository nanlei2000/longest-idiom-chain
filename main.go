package main

import (
	"fmt"
	"os"
	"strconv"

	"log"

	f "github.com/nanlei2000/longest-idiom-chain/findchain"
	"github.com/urfave/cli/v2"
)

func main() {

	var maxLoopCount string
	app := &cli.App{
		Name:  "findchain",
		Usage: "find the longest idiom chain",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "mlc",
				Value:       "100000",
				Usage:       "max loop count for dfs",
				Destination: &maxLoopCount,
			},
		},
		Action: func(c *cli.Context) error {
			idiom := "文武双全"
			if c.NArg() > 0 {
				idiom = c.Args().Get(0)
			}
			graph := f.ReadGraph()
			loopCount, _ := strconv.ParseInt(maxLoopCount, 10, 64)
			idToGraphItemMap := f.MakeIDToGraphItemMap(graph)
			wordToGraphItemMap := f.MakeWordToGraphItemMap(graph)
			wordID := wordToGraphItemMap[idiom].ID
			longest := f.FindLongestChain(wordID, idToGraphItemMap, loopCount)
			words := f.MapIDtoIdiom(longest, idToGraphItemMap)
			fmt.Printf("%s\n", words)
			println(len(longest))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
