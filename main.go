package main

import (
	"flag"
	"log"
	"eagleeye-pushgateway/g"
	"os"
	"os/signal"
	"syscall"
)

func start_signal(pid int, cfg *g.GlobalConfig) {
	sigs := make(chan os.Signal, 1)
	log.Println(pid, "register signal notify")
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	for {
		s := <-sigs
		log.Println("recv", s)
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


	start_signal(os.Getpid(), g.Config())
}
