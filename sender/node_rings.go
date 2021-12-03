package sender

import (
	"eagleeye-pushgateway/utils"
	"eagleeye-pushgateway/g"
	"github.com/toolkits/consistent/rings"
)

func initNodeRings() {
	cfg := g.Config()
	InfluxNodeRing = rings.NewConsistentHashNodesRing(int32(cfg.InfluxDB.Replicas), utils.KeysOfMap(cfg.InfluxDB.Cluster))
}
