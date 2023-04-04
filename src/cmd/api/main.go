package main

import (
	"ben/haute/common/logz"
	"ben/haute/config"
	"ben/haute/servers"
	"flag"
	"os"
)

var Version = "dev"
var log = logz.LOGZ

func main() {
	var stageflag = flag.String("stage", "dev", "stage could be dev, preprod or prod")

	flag.Parse()
	if v, ok := os.LookupEnv("STAGE"); ok {
		stageflag = &v
	} else {
		os.Setenv("STAGE", *stageflag)
	}

	log.Infof("Starting Admin API version %s", Version)

	cfg := config.Instance
	servers.ServePost(cfg.Server.Port, "/", false)
}
