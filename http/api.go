package http

import (
	cmodel "eagleeye-pushgateway/model"
	"eagleeye-pushgateway/utils"
	"eagleeye-pushgateway/sender"
	influx "github.com/influxdata/influxdb/client/v2"
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"time"
	"log"
)

func api_push_datapoints(rw http.ResponseWriter, req *http.Request) {
	if req.ContentLength == 0 {
		http.Error(rw, "blank body", http.StatusBadRequest)
		return
	}
	decoder := json.NewDecoder(req.Body)
	var metrics []*cmodel.MetricValue
	err := decoder.Decode(&metrics)
	if err != nil {
		http.Error(rw, "decode error", http.StatusBadRequest)
		return
	}
	RecvMetricValues(metrics)
	RenderDataJson(rw, nil)
}

func RecvMetricValues(args []*cmodel.MetricValue){
	items := []*cmodel.MetaData{}
	for _ ,v := range args{
		fv := &cmodel.MetaData{
			Metric:     v.Metric,
			Tags:       utils.DictedTagstring(v.Tags),
			Timestamp:  time.Unix(v.Timestamp, 0),
			Value: 		v.Value,
		}
		items = append(items, fv)

		//
		valid := true
		var vv float64
		var err error
		switch cv := v.Value.(type) {
		case string:
			vv, err = strconv.ParseFloat(cv, 64)
			if err != nil {
				valid = false
			}
		case float64:
			vv = cv
		case int64:
			vv = float64(cv)
		default:
			valid = false
		}
		if !valid {
			continue
		}
		if math.IsNaN(vv) || math.IsInf(vv, 0) {
			continue
		}
		//构造influxb的数据格式
		p, err := influx.NewPoint(
			v.Metric,
			utils.DictedTagstring(v.Tags),
			map[string]interface{}{"value": vv},
			time.Unix(v.Timestamp, 0),
		)
		if err != nil {
			continue
		}

		//计算该分配到哪个节点
		//node: node-01
		node, err := sender.InfluxNodeRing.GetNode(utils.PK(v.Metric, utils.DictedTagstring(v.Tags)))
		if err != nil {
			log.Println("E:", err)
			continue
		}
		Q := sender.InfluxQueues[node]
		//写入缓存队列
		Q.PushFront(p)

	}
}

func configApiRoutes() {
	http.HandleFunc("/api/push", api_push_datapoints)
}
