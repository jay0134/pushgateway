package sender

import (
	rings "github.com/toolkits/consistent/rings"
	nlist "github.com/toolkits/container/list"
	"log"
)

// 服务节点的一致性哈希环
// pk -> node
var (
	InfluxNodeRing *rings.ConsistentHashNodeRing
)

// 发送缓存队列
// node -> queue_of_data
var (
	InfluxQueues = make(map[string]*nlist.SafeListLimited)
)

const (
	//缓存能存的最大吞吐量
	DefaultSendQueueMaxSize = 102400 //10.24w
)



func Start() {
	initInfluxClient()
	initSendQueues()
	initNodeRings()
	startSendTasks()
	log.Println("send.Start, ok")
}