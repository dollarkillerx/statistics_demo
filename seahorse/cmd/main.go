package main

import (
	"seahorse/internal/api"
	"seahorse/internal/conf"
	"seahorse/internal/storage"
)

// 数据量太大了就不导入了

func main() {
	conf := conf.LoadConf()
	storage := storage.New(conf)

	server := api.NewApiServer(conf, storage)
	server.Start()
}
