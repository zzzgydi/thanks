package main

import (
	"github.com/zzzgydi/thanks/common/config"
	"github.com/zzzgydi/thanks/common/initializer"
	"github.com/zzzgydi/thanks/common/logger"
	"github.com/zzzgydi/thanks/router"
)

func main() {
	rootDir := config.GetRootDir()
	logger.InitLogger(rootDir)
	config.InitConfig()
	initializer.InitInitializer()
	router.InitHttpServer()
}
