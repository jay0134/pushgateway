package model

//自主上报的结构体
type CustomMetricValue struct {
	Metric    string      `json:"metric"`
	Tags      string      `json:"tags"`
	Timestamp int64       `json:"timestamp"`
	Value     interface{} `json:"value"`
}

