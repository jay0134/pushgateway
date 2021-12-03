package sender


import (
	"eagleeye-pushgateway/g"
	nlist "github.com/toolkits/container/list"
)

func initSendQueues() {
	cfg := g.Config()
	for node := range cfg.InfluxDB.Cluster {
		Q := nlist.NewSafeListLimited(DefaultSendQueueMaxSize)
		InfluxQueues[node] = Q
	}
}
