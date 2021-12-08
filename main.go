package main

import (
	"context"
	"flag"
	"log"
	"eagleeye-pushgateway/g"
	"eagleeye-pushgateway/http"
	"eagleeye-pushgateway/sender"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func start_signal(pid int, cfg *g.GlobalConfig) {
	sigs := make(chan os.Signal, 1)
	log.Println(pid, "register signal notify")
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	for {
		s := <-sigs
		log.Println("recv", s)
		//关闭http服务
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		http.Srv.Shutdown(ctx)
		cancel()
		// 判断缓存队列是否都已处理完毕
		empty := true
		for {
			for node := range cfg.InfluxDB.Cluster {
				queue := sender.InfluxQueues[node]
				if queue.Len() != 0{
					empty = false
					break
				}
			}
			if empty == true{
				break
			}else{
				time.Sleep(1* time.Second)
			}
		}
		os.Exit(0)
	}
}

func main() {
	cfg := flag.String("c", "cfg.json", "specify config file")
	flag.Parse()
	g.ParseConfig(*cfg)
	if g.Config().Debug {
		g.InitLog("debug")
	} else {
		g.InitLog("info")
	}

	sender.Start()
	go http.Start()
	start_signal(os.Getpid(), g.Config())
}
