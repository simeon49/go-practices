##gossip

先理解一下gossip协议：在一个有界网络中，每个节点都随机地与其他节点通信，经过一番杂乱无章的通信，最终所有节点的状态都会达成一致。每个节点可能知道所有其他节点，也可能仅知道几个邻居节点，只要这些节可以通过网络连通，最终他们的状态都是一致的，当然这也是疫情传播的特点。
简单的描述下这个协议，首先要传播谣言就要有种子节点。种子节点每秒都会随机向其他节点发送自己所拥有的节点列表，以及需要传播的消息。任何新加入的节点，就在这种传播方式下很快地被全网所知道。这个协议的神奇就在于它从设计开始就没想到信息一定要传递给所有的节点，但是随着时间的增长，在最终的某一时刻，全网会得到相同的信息。当然这个时刻可能仅仅存在于理论，永远不可达。

参考: [https://zhuanlan.zhihu.com/p/41228196](https://zhuanlan.zhihu.com/p/41228196)

#### 使用方法
这里通过一个简单的http服务查询和插入数据，找两台机器，第一台执行
```bash
$ go run ./main.go
```
会生成gossip监听的服务ip和端口
使用上面的ip和端口在第二台执行
```bash
$ go run ./main.go --merbers=xxx.xxx.xxx.xxx:ppp
```
那么一个gossip的网络就搭建完成了。

```bash
# add
curl "http://localhost:4001/add?key=foo&val=bar"

# get
curl "http://另一台机器:4001/get?key=foo"

# delete
curl "http://localhost:4001/del?key=foo"
```
