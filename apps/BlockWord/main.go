package main

import (
	"github.com/mangenotwork/extras/apps/BlockWord/engine"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/utils"
	"log"
)

func main(){

	// 打印logo
	log.Println(utils.Logo)
	log.Println("Starting block word http server")

	conf.InitConf()

	engine.StartJobSrc()
	engine.StartRpcSrc()
	engine.StartHttpSrc()
}
