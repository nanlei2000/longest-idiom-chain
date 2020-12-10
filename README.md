## Longest Idiom Chain

- Use `./files/idiom.json` to make a graph and store in `./files/graph.json`
```sh
go run makegraph/makegraph.go
```

- Find the longest path in the graph
```sh
go run findchain/findchain.go
```

In this [page](http://www.jielongdaquan.com/phrase/chengyujielong.aspx?pageIndex=1), it declared he found the longest idiom chain. It's length is `9210`, and start with `文武双全`.While i had found a chain with this leading idiom have length of `9321`!