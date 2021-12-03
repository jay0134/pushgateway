package sender

import (
	"eagleeye-pushgateway/g"
	"eagleeye-pushgateway/utils"
	influx "github.com/influxdata/influxdb/client/v2"
	"os"
	"time"
)

var (
	InfluxClientMap = make(map[string]*Client)
)


type Client struct {
	client          influx.Client
	database        string
	retentionPolicy string
}


func NewClient(conf influx.HTTPConfig, db string, rp string) *Client {
	c, err := influx.NewHTTPClient(conf)
	// Currently influx.NewClient() *should* never return an error.
	if err != nil {
		os.Exit(1)
	}

	return &Client{
		client:          c,
		database:        db,
		retentionPolicy: rp,
	}
}

//初始化集群中各节点的influxdb client
func initInfluxClient(){
	cfg := g.Config()
	var timeout int
	timeout = utils.TimeStringToSec(cfg.InfluxDB.Timeout)
	for node, address := range cfg.InfluxDB.Cluster {
		conf := influx.HTTPConfig{
			Addr:     address,
			Username: cfg.InfluxDB.Username,
			Password: cfg.InfluxDB.Password,
			Timeout:  time.Duration(timeout) * time.Second,
		}
		c := NewClient(conf, cfg.InfluxDB.DB, cfg.InfluxDB.RetentionPolicy)
		InfluxClientMap[node] = c
	}
	return
}
