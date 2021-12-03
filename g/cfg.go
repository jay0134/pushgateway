package g

import (
	"encoding/json"
	"github.com/toolkits/file"
	"log"
	"sync/atomic"
	"unsafe"
)

type HttpConfig struct {
	Listen  string `json:"listen"`
}

// CLUSTER NODE
type ClusterNode struct {
	Addrs []string `json:"addrs"`
}

type InfluxDBConfig struct {
	DB          	string 					`json:"db"`
	Username    	string 					`json:"username"`
	Password   	 	string					`json:"password"`
	Timeout     	string  				`json:"timeout"`
	RetentionPolicy string 					`json:"retention-policy"`
	Replicas   		int                     `json:"replicas"`
	Batch       	int                     `json:"batch"`
	Cluster     	map[string]string	    `json:"cluster"`
	ClusterList 	map[string]*ClusterNode `json:"clusterList"`
}



type GlobalConfig struct {
	Debug          bool             `json:"debug"`
	Http           *HttpConfig      `json:"http"`
	InfluxDB       *InfluxDBConfig  `json:"influxdb"`
}

var (
	ConfigFile string
	ptr        unsafe.Pointer
)

func Config() *GlobalConfig {
	return (*GlobalConfig)(atomic.LoadPointer(&ptr))
}




func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("config file not specified: use -c $filename")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file specified not found:", cfg)
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file", cfg, "error:", err.Error())
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file", cfg, "error:", err.Error())
	}

	// set config
	atomic.StorePointer(&ptr, unsafe.Pointer(&c))

	log.Println("g.ParseConfig ok, file", cfg)
}
