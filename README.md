## Longest Idiom Chain

尝试找到最长的成语接龙！

```sh
// 测试迭代一百万次 
go run . -mlc 1_000_000
```
输出
```
[info] head idiom: 文武双全, max loop count: 1000000
chain: [文武双全 全功尽弃 ...
length: 9321
dfs took: 7520ms
```

In this [page](http://www.jielongdaquan.com/phrase/chengyujielong.aspx?pageIndex=1), it declared he found the longest idiom chain. It's length is `9210`, and start with `文武双全`.While i had found a chain with this leading idiom have length of [`9662`](./files/文武双全-9662.json)!
