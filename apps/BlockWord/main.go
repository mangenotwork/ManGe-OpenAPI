package main

import (
	"github.com/mangenotwork/extras/apps/BlockWord/engine"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

func main(){
	conf.InitConf()
	logger.InitLogger()

	logger.Info(utils.Logo)
	logger.Info("Starting block word http server")

	engine.StartJobServer()

	if conf.Arg.HttpServer.Open {
		engine.StartHttp()
	}

	if conf.Arg.GrpcServer.Open {
		engine.StartRpcServer()
	}

	select {}
}
