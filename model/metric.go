package model

import (
	"eagleeye-pushgateway/utils"
	"time"
)

type MetricValue struct {
	Metric    string      `json:"metric"`
	Tags      string      `json:"tags"`
	Timestamp int64       `json:"timestamp"`
	Value     interface{} `json:"value"`
}



type MetaData struct {
	Metric      string            `json:"metric"`
	Tags        map[string]string `json:"tags"`
	Timestamp   time.Time         `json:"timestamp"`
	Value       interface{}       `json:"value"`
}

func (t *MetaData) PK() string {
	return utils.PK(t.Metric, t.Tags)
}


type InluxItem struct {
	Metric      string            `json:"metric"`
	Tags        map[string]string `json:"tags"`
	Timestamp   time.Time         `json:"timestamp"`
	Value       interface{}       `json:"value"`
}