package g

import (
	"encoding/json"
	"log"
	"sync/atomic"
	"unsafe"
	"github.com/toolkits/file"
)

type HttpConfig struct {
	Address  string `json:"address"`
}

type InfluxDBConfig struct {
	Address   string `json:"address"`
	DB        string `json:"db"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Timeout   string `json:"timeout"`

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
