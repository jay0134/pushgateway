package sender

import (
	"eagleeye-pushgateway/g"
	influx "github.com/influxdata/influxdb/client/v2"
	"github.com/toolkits/container/list"
	"time"
)

// send
const (
	DefaultSendTaskSleepInterval = time.Millisecond * 50 //默认睡眠间隔为50ms
)


func startSendTasks() {
	cfg := g.Config()
	// init send go-routines，向每个judge节点发送数据
	for node := range cfg.InfluxDB.Cluster {
		queue := InfluxQueues[node]
		go forward2InluxTask(queue, node)
	}
}

// 将缓存中的数据发送到influxdb
func forward2InluxTask(Q *list.SafeListLimited, node string) {
	batch := g.Config().InfluxDB.Batch // 一次发送,最多batch条数据

	for {
		items := Q.PopBackBy(batch)		//一次性向队列尾部(右边)取出batch个数据
		count := len(items)
		if count == 0 {
			time.Sleep(DefaultSendTaskSleepInterval)
			continue
		}

		infuxItems := make([]*influx.Point, count)
		for i := 0; i < count; i++ {
			infuxItems[i] = items[i].(*influx.Point)
		}

		//写入influxdb
		bps, err := influx.NewBatchPoints(influx.BatchPointsConfig{
			Precision:       "ms",
			Database: g.Config().InfluxDB.DB,
			RetentionPolicy: g.Config().InfluxDB.RetentionPolicy,
		})
		if err != nil {
		}
		bps.AddPoints(infuxItems)
		if err := InfluxClientMap[node].client.Write(bps); err != nil{

		}
	}
}